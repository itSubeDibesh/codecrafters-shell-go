package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isExecutablePath(cmd string) (string, bool) {
	osPath := os.Getenv("PATH")
	paths := strings.Split(osPath, ":")
	for _, path := range paths {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", path, cmd)); err == nil {
			return path,true
		}
	}
	return "", false
}

func isBuiltin(name string) bool {
	switch name {
		case "exit","echo","type":
			return true
		default:
			return false
	}
}

func handelExecType(cmd string, input string){
	response := ""
	for _, word := range strings.Fields(input){
		if isBuiltin(word){
			response = fmt.Sprintf("%s is a shell builtin\n", word)
		}else if path, ok := isExecutablePath(word);ok {
			response = fmt.Sprintf("%s is %s/%s\n", word, path, word)
		}else{
			response = fmt.Sprintf("%s not found\n", word)
		}
	}	
	fmt.Printf("%s", response)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")
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
