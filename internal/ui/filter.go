package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterModel struct {
	TextInput textinput.Model
}

func NewFilter() FilterModel {
	ti := textinput.New()
	ti.Placeholder = "Buscar por puerto, proceso o PID..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return FilterModel{
		TextInput: ti,
	}
}

func (m FilterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FilterModel) Update(msg tea.Msg) (FilterModel, tea.Cmd) {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m FilterModel) View() string {
	return "Filtrar: " + m.TextInput.View()
}
