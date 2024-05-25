package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/exec"
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

func handelExecType(cmd string, args []string){
	response := ""
	for _, word := range args{
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

func handelExec(cmd string, args []string){
	exec := exec.Command(cmd, args...)
	exec.Stdout = os.Stdout
	exec.Stderr = os.Stderr

	err := exec.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func handelInput(input string)(cmd string, args []string){
	args = strings.Fields(input)
	cmd = args[0]
	args = args[1:]
	return cmd, args
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")
	const prefixCommands = "echo type"
	reader := bufio.NewReader(os.Stdin)
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputCopy := input
		cmd, args := handelInput(input)

		for _, word := range strings.Fields(prefixCommands){
			if strings.HasPrefix(input, fmt.Sprintf("%s ", word)){
				input = word
				break;
			}
		}

		switch input {
			case "exit 0":
				os.Exit(0)
			case "echo":
				fmt.Printf("%s\n", strings.TrimPrefix(inputCopy, "echo "))
			case "type":
				handelExecType(cmd, args)
			default:
				handelExec(cmd, args)
		}

	}
}
