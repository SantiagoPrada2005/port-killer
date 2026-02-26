package process

import (
	"fmt"
	"os/exec"
)

// KillProcess terminates a process given its PID and a signal (e.g. "SIGTERM", "SIGKILL")
func KillProcess(pid, signal string) error {
	sigFlag := "-15"
	if signal == "SIGKILL" {
		sigFlag = "-9"
	}

	cmd := exec.Command("kill", sigFlag, pid)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(output))
	}
	return nil
}
