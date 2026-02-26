package ui

import (
	"fmt"
	"strings"

	"port-killer/internal/ports"
	"port-killer/internal/process"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type State int

const (
	StateTable State = iota
	StateDetail
	StateFilter
	StateConfirm
)

type MainModel struct {
	state       State
	table       table.Model
	filter      FilterModel
	confirmForm *huh.Form

	err         error
	feedbackMsg string

	allPorts     []ports.ProcessPort
	visiblePorts []ports.ProcessPort
	selectedPort ports.ProcessPort
}

func InitialModel() MainModel {
	columns := []table.Column{
		{Title: "PORT", Width: 8},
		{Title: "PROTOCOL", Width: 10},
		{Title: "PID", Width: 8},
		{Title: "PROCESS", Width: 15},
		{Title: "STATUS", Width: 18},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := MainModel{
		state:  StateTable,
		table:  t,
		filter: NewFilter(),
	}
	m.loadPorts()
	return m
}

func (m *MainModel) loadPorts() {
	p, err := ports.Scan()
	if err != nil {
		m.err = err
		return
	}
	m.allPorts = p
	m.applyFilter()
}

func (m *MainModel) applyFilter() {
	term := strings.ToLower(m.filter.TextInput.Value())
	if term == "" {
		m.visiblePorts = m.allPorts
	} else {
		var filtered []ports.ProcessPort
		for _, p := range m.allPorts {
			if strings.Contains(strings.ToLower(p.Port), term) ||
				strings.Contains(strings.ToLower(p.Process), term) ||
				strings.Contains(strings.ToLower(p.PID), term) {
				filtered = append(filtered, p)
			}
		}
		m.visiblePorts = filtered
	}
	m.updateTable()
}

func (m *MainModel) updateTable() {
	var rows []table.Row
	for _, p := range m.visiblePorts {
		rows = append(rows, table.Row{
			p.Port,
			p.Protocol,
			p.PID,
			p.Process,
			RenderStatus(p.Status),
		})
	}
	m.table.SetRows(rows)
}

func (m *MainModel) getSelectedPort() *ports.ProcessPort {
	if len(m.visiblePorts) == 0 {
		return nil
	}
	idx := m.table.Cursor()
	if idx >= 0 && idx < len(m.visiblePorts) {
		return &m.visiblePorts[idx]
	}
	return nil
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.filter.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch m.state {
	case StateTable:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			// Limpiar feedback en nueva acción
			m.feedbackMsg = ""

			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "r":
				m.loadPorts()
				return m, nil
			case "enter":
				if p := m.getSelectedPort(); p != nil {
					m.selectedPort = *p
					m.state = StateDetail
				}
				return m, nil
			case "k", "K":
				if p := m.getSelectedPort(); p != nil {
					m.selectedPort = *p
					m.confirmForm = NewConfirmForm(m.selectedPort)
					m.confirmForm.Init()
					m.state = StateConfirm
				}
				return m, nil
			case "f", "F":
				m.state = StateFilter
				m.filter.TextInput.Focus()
				return m, textinput.Blink
			}
		}

		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)

	case StateFilter:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter, tea.KeyEsc:
				m.state = StateTable
				m.applyFilter()
				return m, nil
			}
		}
		m.filter, cmd = m.filter.Update(msg)
		cmds = append(cmds, cmd)
		m.applyFilter() // Filtrar en vivo

	case StateDetail:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "q":
				m.state = StateTable
			case "k", "K":
				m.confirmForm = NewConfirmForm(m.selectedPort)
				m.confirmForm.Init()
				m.state = StateConfirm
			}
		}

	case StateConfirm:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "esc" {
				m.state = StateTable
				return m, nil
			}
		}

		form, formCmd := m.confirmForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.confirmForm = f
			cmds = append(cmds, formCmd)

			if m.confirmForm.State == huh.StateCompleted {
				signal := m.confirmForm.GetString("signal")
				err := process.KillProcess(m.selectedPort.PID, signal)
				if err != nil {
					m.feedbackMsg = fmt.Sprintf("✗ Error al terminar proceso: %s", err)
				} else {
					m.feedbackMsg = fmt.Sprintf("✓ Proceso %q (PID %s) terminado correctamente.", m.selectedPort.Process, m.selectedPort.PID)
					m.loadPorts() // Recargar puertos
				}
				m.state = StateTable
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if m.err != nil {
		return "Error loading ports: " + m.err.Error() + "\n"
	}

	var content string

	switch m.state {
	case StateTable:
		header := TitleStyle.Render("⚡ Port Killer") + lipgloss.NewStyle().Foreground(NormalColor).Render("   Refresh: r  Quit: q")
		tview := m.table.View()
		footer := HelpStyle.Render("↑/↓: Navegar  Enter: Detalles  K: Matar  F: Filtrar")

		view := lipgloss.JoinVertical(lipgloss.Left, header, ContainerStyle.Render(tview), footer)

		if m.feedbackMsg != "" {
			color := ListeningColor
			if strings.HasPrefix(m.feedbackMsg, "✗") {
				color = CloseWaitColor
			}
			fbStyle := lipgloss.NewStyle().Foreground(color).Margin(1, 0, 0, 0)
			view = lipgloss.JoinVertical(lipgloss.Left, view, fbStyle.Render(m.feedbackMsg))
		}
		content = view

	case StateFilter:
		header := TitleStyle.Render("⚡ Port Killer - Filtrando")
		tview := m.table.View()
		filterView := lipgloss.NewStyle().Margin(1, 0).Render(m.filter.View())

		content = lipgloss.JoinVertical(lipgloss.Left, header, filterView, ContainerStyle.Render(tview))

	case StateDetail:
		content = RenderDetail(m.selectedPort)

	case StateConfirm:
		content = BaseStyle.Render(m.confirmForm.View())
	}

	return BaseStyle.Render(content)
}
