package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Básico
	HighlightColor = lipgloss.Color("205") // Rosa brillante
	NormalColor    = lipgloss.Color("240") // Gris

	// Estados
	ListeningColor   = lipgloss.Color("42")  // Verde
	EstablishedColor = lipgloss.Color("220") // Amarillo
	CloseWaitColor   = lipgloss.Color("196") // Rojo

	// Bordes & Padding
	BaseStyle = lipgloss.NewStyle().
			Margin(1, 0)

	ContainerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(NormalColor)

	TitleStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Bold(true).
			Margin(0, 1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(NormalColor).
			Margin(1, 0, 0, 0)

	StatusListening   = lipgloss.NewStyle().Foreground(ListeningColor).Render("🟢 LISTENING")
	StatusEstablished = lipgloss.NewStyle().Foreground(EstablishedColor).Render("🟡 ESTABLISHED")
	StatusCloseWait   = lipgloss.NewStyle().Foreground(CloseWaitColor).Render("🔴 CLOSE_WAIT")
)

func RenderStatus(status string) string {
	switch status {
	case "LISTENING":
		return StatusListening
	case "ESTABLISHED":
		return StatusEstablished
	case "CLOSE_WAIT":
		return StatusCloseWait
	default:
		return status
	}
}
