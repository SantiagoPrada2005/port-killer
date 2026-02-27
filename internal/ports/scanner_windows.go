//go:build windows

package ports

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func Scan() ([]ProcessPort, error) {
	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		if len(output) == 0 {
			return nil, err
		}
	}

	var results []ProcessPort
	scanner := bufio.NewScanner(bytes.NewReader(output))

	// Skip 4 lines of header commonly found in Windows netstat
	for i := 0; i < 4; i++ {
		if !scanner.Scan() {
			return results, nil
		}
	}

	pidProcessMap := getProcessMap()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}

		protocol := fields[0]
		localAddress := fields[1]

		var status, pid string

		if protocol == "TCP" {
			if len(fields) < 5 {
				continue
			}
			status = fields[3]
			pid = fields[4]
		} else if protocol == "UDP" {
			status = "" // UDP doesn't have a state in Windows netstat
			pid = fields[3]
		} else {
			continue
		}

		port := parsePortNetstat(localAddress)
		if port == "" || port == "0" {
			continue
		}

		if status == "LISTENING" {
			// keep AS IS, UI handles it
		}

		processName := "Unknown"
		if name, ok := pidProcessMap[pid]; ok {
			processName = name
		}

		results = append(results, ProcessPort{
			Process:  processName,
			PID:      pid,
			User:     "N/A",
			Protocol: protocol,
			Port:     port,
			Status:   status,
		})
	}

	return results, nil
}

func parsePortNetstat(address string) string {
	lastColon := strings.LastIndex(address, ":")
	if lastColon != -1 && lastColon+1 < len(address) {
		return address[lastColon+1:]
	}
	return ""
}

func getProcessMap() map[string]string {
	processMap := make(map[string]string)

	cmd := exec.Command("tasklist", "/FO", "CSV", "/NH")
	output, err := cmd.Output()
	if err != nil {
		return processMap
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))

	// CSV format is like: "System Idle Process","0","Services","0","8 K"
	re := regexp.MustCompile(`"([^"]+)","([^"]+)"`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 3 {
			processName := matches[1]
			pid := matches[2]
			processMap[pid] = processName
		}
	}

	return processMap
}
