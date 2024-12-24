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
	if !exists {
		return fmt.Errorf("%s: not found", cmd.args[0])
	}

	fmt.Printf("%s is a shell builtin\n", cmd.args[0])
	return nil
}
func Eval(cmd *Command) error {
	fMaps := GetFuncMap()
	if fn, exists := fMaps[cmd.cmd]; exists {
		err := fn(cmd)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("%s: command not found\n", cmd.cmd)
	}

	return nil
}
