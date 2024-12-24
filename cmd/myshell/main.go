package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		ansString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s: command not found\n", ansString[:len(ansString)-1])
	}

}
