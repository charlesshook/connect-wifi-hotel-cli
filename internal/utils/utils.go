package utils

import (
	"os/exec"
	"strings"
)

func RunCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(output)), err
}
