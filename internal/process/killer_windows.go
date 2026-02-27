//go:build windows

package process

import (
	"fmt"
	"os/exec"
)

// KillProcess terminates a process given its PID and a signal (e.g. "SIGTERM", "SIGKILL")
func KillProcess(pid, signal string) error {
	args := []string{"/PID", pid}
	if signal == "SIGKILL" {
		args = append(args, "/F")
	}

	cmd := exec.Command("taskkill", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(output))
	}
	return nil
}
