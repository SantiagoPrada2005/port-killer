package ui

import (
	"fmt"
	"port-killer/internal/ports"

	"github.com/charmbracelet/lipgloss"
)

func RenderDetail(p ports.ProcessPort) string {
	title := TitleStyle.Render("Detalles del Proceso")

	content := fmt.Sprintf(
		"Proceso   : %s\n"+
			"PID       : %s\n"+
			"Puerto    : %s\n"+
			"Protocolo : %s\n"+
			"Estado    : %s\n"+
			"Usuario   : %s\n",
		p.Process, p.PID, p.Port, p.Protocol, RenderStatus(p.Status), p.User)

	box := ContainerStyle.Render(content)
	footer := HelpStyle.Render("[ K ] Matar    [ Esc ] Volver")

	return BaseStyle.Render(lipgloss.JoinVertical(lipgloss.Left, title, box, footer))
}
