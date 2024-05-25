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
		switch strings.TrimSpace(cmd) {
			case "exit 0":
				os.Exit(0)
			default:
				fmt.Printf("%s: command not found\n", strings.TrimSpace(cmd))
		}
	}
}
