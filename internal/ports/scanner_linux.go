//go:build linux

package ports

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func Scan() ([]ProcessPort, error) {
	cmd := exec.Command("ss", "-tulnp")
	output, err := cmd.Output()
	if err != nil {
		if len(output) == 0 {
			return nil, err
		}
	}

	var results []ProcessPort
	scanner := bufio.NewScanner(bytes.NewReader(output))

	if scanner.Scan() {
		_ = scanner.Text() // Skip header
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue // Skip invalid lines
		}

		protocol := strings.ToUpper(fields[0])

		var status string
		var localAddressIdx int

		switch protocol {
		case "TCP", "TCP6":
			status = fields[1]
			localAddressIdx = 4
		case "UDP", "UDP6":
			status = fields[1] // ss shows state for UDP too
			localAddressIdx = 4
		default:
			continue
		}

		// Find users column
		processInfoStr := ""
		for i := localAddressIdx + 1; i < len(fields); i++ {
			if strings.HasPrefix(fields[i], "users:(") {
				processInfoStr = fields[i]
				break
			}
		}

		if processInfoStr == "" {
			continue // Can't identify process, maybe permission issue
		}

		localAddress := fields[localAddressIdx]
		port := parsePortSS(localAddress)
		if port == "" || port == "*" {
			continue
		}

		pid, processName := parseProcessInfoSS(processInfoStr)

		switch protocol {
		case "TCP6":
			protocol = "TCP"
		case "UDP6":
			protocol = "UDP"
		}

		if status == "LISTEN" {
			status = "LISTENING"
		}

		results = append(results, ProcessPort{
			Process:  processName,
			PID:      pid,
			User:     "N/A", // ss doesn't easily show user without extra flags like -e, keeping it simple
			Protocol: protocol,
			Port:     port,
			Status:   status,
		})
	}

	return results, nil
}

func parsePortSS(address string) string {
	lastColon := strings.LastIndex(address, ":")
	if lastColon != -1 && lastColon+1 < len(address) {
		return address[lastColon+1:]
	}
	return ""
}

func parseProcessInfoSS(info string) (string, string) {
	// Format: users:(("process",pid=123,fd=4))
	info = strings.TrimPrefix(info, "users:((\"")
	info = strings.TrimSuffix(info, "))")

	parts := strings.Split(info, "\",pid=")
	if len(parts) == 2 {
		processName := parts[0]
		pidFd := parts[1]
		pid := strings.Split(pidFd, ",")[0]
		return pid, processName
	}
	return "", ""
}
