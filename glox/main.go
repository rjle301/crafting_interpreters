package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println("Hello, Glox!")
	if len(os.Args) > 2 {
		fmt.Println("Usage: go run glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		fmt.Println("runFile(os.Args[1])")
		runFile(os.Args[1])
	} else {
		fmt.Println("runPrompt()")
		runPrompt()
	}
}

func runFile(filepath string) {
	bytes, err := os.ReadFile(path.Clean(filepath))
	if err != nil {
		os.Exit(65)
	}
	fmt.Println(bytes)
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}
		fmt.Println(scanner.Text())
	}
}

// func run(line string) {
// 	fmt.Print(line)
// }
