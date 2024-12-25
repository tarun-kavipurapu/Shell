package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		ansString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		commands := strings.Fields(ansString)
		cmd := &Command{
			cmd:  commands[0],
			args: commands[1:],
		}
		err = Eval(cmd)
		if err != nil {
			fmt.Print(err, "\n")
		}

	}

}
