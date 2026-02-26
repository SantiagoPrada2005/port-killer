package ui

import (
	"fmt"
	"port-killer/internal/ports"

	"github.com/charmbracelet/huh"
)

func NewConfirmForm(p ports.ProcessPort) *huh.Form {
	title := fmt.Sprintf("¿Matar el proceso %q (PID %s)?\nPuerto: %s", p.Process, p.PID, p.Port)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("signal").
				Title(title).
				Description("Selecciona el tipo de señal a enviar:").
				Options(
					huh.NewOption("SIGTERM (Terminación elegante)", "SIGTERM"),
					huh.NewOption("SIGKILL (Forzar cierre)", "SIGKILL"),
				),
		),
	)

	return form
}
