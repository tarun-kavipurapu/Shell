package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Command structure
type Command struct {
	cmd  string
	args []string
}

// Employing a singleton pattern for this
var (
	funcMap map[string]func(cmd *Command) error
	once    sync.Once
)

func GetFuncMap() map[string]func(cmd *Command) error {
	once.Do(func() {
		funcMap = map[string]func(cmd *Command) error{
			"exit": evalExit,
			"echo": evalEcho,
			"type": evalType,
			"pwd":  evalPwd,
			"cd":   evalCd,
		}
	})

	return funcMap
}

func evalExit(cmd *Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("invalid command args for exit")
	}
	exitcode, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid exit code conversion for exit command")
	}
	os.Exit(exitcode)
	return nil
}

func evalEcho(cmd *Command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("invalid command args for echo")
	}
	words := strings.Join(cmd.args, " ")
	fmt.Println(words)
	return nil
}
func evalType(cmd *Command) error {
	fMaps := GetFuncMap()
	if len(cmd.args) > 1 {
		return fmt.Errorf("%s:argument error", cmd.cmd)
	}
	_, exists := fMaps[cmd.args[0]]
	if exists {
		fmt.Printf("%s is a shell builtin\n", cmd.args[0])
		return nil
	}

	path, err := getPath(cmd)
	if path != "" {
		fmt.Println(path)
		return nil
	}
	if err != nil {
		return err
	}

	return fmt.Errorf("%s: not found", cmd.args[0])
}
func evalPwd(cmd *Command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("too many args")
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}

func evalCd(cmd *Command) error {

	newPath := cmd.args[0]
	if newPath == "~" {
		HOME := os.Getenv("HOME")
		newPath = HOME

	}
	err := os.Chdir(newPath)
	if err != nil {
		return fmt.Errorf("cd: /non-existing-directory: No such file or directory")
	}

	return nil

}

/*
if the command is not in the map then it may be a program?..
then you can check to getPath() if the path is not empty then you can just exec the binary in that path
*/
func Eval(cmd *Command) error {
	fMaps := GetFuncMap()
	if fn, exists := fMaps[cmd.cmd]; exists {
		err := fn(cmd)
		if err != nil {
			return err
		}
	} else {

		err := executeCommand(cmd)
		if err != nil {
			return err
		}

	}

	return nil
}
