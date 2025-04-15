package logic

import (
	"os/exec"
	"strings"
)

func PickFile() (string, error) {
	cmd := exec.Command("osascript", "-e", `POSIX path of (choose file with prompt "Select file to shred")`)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
