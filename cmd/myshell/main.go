package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handelExecType(cmd string, input string){
	response := ""
	for _, word := range strings.Fields(input){
		switch word {
			case "exit","echo":
				response = fmt.Sprintf("%s is a shell builtin\n", word)
			case "cat":
				response = fmt.Sprintf("%s is /bin/%s\n", word, word)
			default:
				response = fmt.Sprintf("%s not found\n", word)
		}
	}
	if input == "type type"{
		response = fmt.Sprintf("%s is a shell builtin\n", cmd)
	}
	fmt.Printf("%s", response)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")
	const builtIns = "echo exit cat"
	const prefixCommands = "echo type"
	reader := bufio.NewReader(os.Stdin)
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		input := cmd

		for _, word := range strings.Fields(prefixCommands){
			if strings.HasPrefix(cmd, fmt.Sprintf("%s ", word)){
				cmd = word
				break;
			}
		}

		switch cmd {
			case "exit 0":
				os.Exit(0)
			case "echo":
				fmt.Printf("%s\n", strings.TrimPrefix(input, "echo "))
			case "type":
				handelExecType(cmd, input)
			default:
				fmt.Printf("%s: command not found\n", cmd)
		}

	}
}
