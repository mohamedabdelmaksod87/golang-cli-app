package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		fmt.Println("echo", cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	//Field func already trims every white space in our string input
	words := strings.Fields(lowered)
	return words
}
