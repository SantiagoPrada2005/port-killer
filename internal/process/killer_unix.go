//go:build !windows

package process

import (
	"os"
	"strconv"
	"syscall"
)

// KillProcess terminates a process given its PID and a signal (e.g. "SIGTERM", "SIGKILL")
func KillProcess(pidStr, signalStr string) error {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return err
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	sig := syscall.SIGTERM
	if signalStr == "SIGKILL" {
		sig = syscall.SIGKILL
	}

	return proc.Signal(sig)
}
