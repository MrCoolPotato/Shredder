package logic

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

func PickFile() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("osascript", "-e", `POSIX path of (choose file with prompt "Select file to shred")`)
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil

	case "windows":
		filePath, err := dialog.File().Title("Select file to shred").Load()
		if err != nil {
			return "", err
		}
		return filePath, nil

	default:
		return "", errors.New("file dialog system not adapted for this OS")
	}
}
