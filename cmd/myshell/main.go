package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

type command struct {
	cmd  string
	args []string
}

func eval(cmd *command) error {

	if cmd.cmd == "exit" {
		if len(cmd.args) > 1 {
			return fmt.Errorf("invalid command args for exit")
		}
		exitcode, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid exit code conversion for exit command")
		}
		os.Exit(exitcode)
	}

	return nil
}

func main() {
	// Uncomment this block to pass the first stage
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		ansString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		commands := strings.Fields(ansString)
		cmd := &command{
			cmd:  commands[0],
			args: commands[1:],
		}
		err = eval(cmd)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%s: command not found\n", ansString[:len(ansString)-1])
	}

}
