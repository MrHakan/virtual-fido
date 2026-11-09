//go:build windows

package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Execute USB IP attach for Windows
func platformUSBIPExec() *exec.Cmd {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	directory := filepath.Join(filepath.Dir(executable), "usbip")
	if _, err := os.Stat(filepath.Join(directory, "usbip.exe")); os.IsNotExist(err) {
		// Keep supporting go run from the source repository.
		directory, err = filepath.Abs(filepath.Join("cmd", "demo", "usbip", "bin"))
		if err != nil {
			panic(err)
		}
	}
	command := exec.Command(filepath.Join(directory, "usbip.exe"), "attach", "-r", "127.0.0.1", "-b", "2-2")
	command.Dir = directory
	return command
}
