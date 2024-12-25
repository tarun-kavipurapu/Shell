package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getPath(cmd *Command) (string, error) {

	executable := cmd.cmd
	if len(cmd.args) > 0 && cmd.args[0] != "" {
		executable = cmd.args[0]
	}

	// Iterate over directories in PATH
	pathString := os.Getenv("PATH")
	paths := strings.Split(pathString, ":")
	for _, path := range paths {
		fp := filepath.Join(path, executable)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
	}
	return "", nil
}
func executeCommand(cmd *Command) error {
	var comm *exec.Cmd

	// Check if cmd.args has elements
	if len(cmd.args) > 0 {
		comm = exec.Command(cmd.cmd, cmd.args...)
	} else {
		comm = exec.Command(cmd.cmd)
	}

	// Set Stdout and Stderr to the same as the current process
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr

	// Run the command
	if err := comm.Run(); err != nil {
		return fmt.Errorf("%s: command not found", cmd.cmd)
	}

	return nil
}
