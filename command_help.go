package main

import "fmt"

func callbackHelp() {
	fmt.Println("Available Commands:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
}
