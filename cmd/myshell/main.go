package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	reader := bufio.NewReader(os.Stdin)
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		input := cmd
		if strings.HasPrefix(cmd, "echo "){
			cmd = "echo"
		}
		switch cmd {
			case "exit 0":
				os.Exit(0)
			case "echo":
				fmt.Printf("%s\n", strings.TrimPrefix(input, "echo "))
			default:
				fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
