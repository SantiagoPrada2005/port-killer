//go:build darwin

package ports

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func Scan() ([]ProcessPort, error) {
	cmd := exec.Command("lsof", "-i", "-P", "-n")
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

		if len(fields) < 9 {
			continue // Skip invalid lines
		}

		command := fields[0]
		pid := fields[1]
		user := fields[2]
		protocol := fields[7]
		nodeName := strings.Join(fields[8:], " ")

		port := parsePort(nodeName)
		if port == "" || port == "*" {
			continue // Could not parse port or port is *, skip
		}
		status := parseStatus(nodeName)

		// Solo agregamos TCP/UDP válidos
		if protocol != "TCP" && protocol != "UDP" {
			continue
		}

		results = append(results, ProcessPort{
			Process:  command,
			PID:      pid,
			User:     user,
			Protocol: protocol,
			Port:     port,
			Status:   status,
		})
	}

	return results, nil
}

func parsePort(nodeName string) string {
	parts := strings.Split(nodeName, " ")
	ipPortPart := parts[0]

	if idx := strings.Index(ipPortPart, "->"); idx != -1 {
		ipPortPart = ipPortPart[:idx]
	}

	lastColon := strings.LastIndex(ipPortPart, ":")
	if lastColon != -1 && lastColon+1 < len(ipPortPart) {
		return ipPortPart[lastColon+1:]
	}
	return ""
}

func parseStatus(nodeName string) string {
	if strings.Contains(nodeName, "(LISTEN)") {
		return "LISTENING"
	} else if strings.Contains(nodeName, "(ESTABLISHED)") {
		return "ESTABLISHED"
	} else if strings.Contains(nodeName, "(CLOSE_WAIT)") {
		return "CLOSE_WAIT"
	}
	return ""
}
