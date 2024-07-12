package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to blnote!")
}

func init() {
	fmt.Println("Welcome to blnote!")
	stdin := bufio.NewReader(os.Stdin)
Infinite:
	for {
		fmt.Println("Please enter a command:")
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error getting input:", err)
			os.Exit(1)
		}
		input = strings.TrimRight(input, " \n")
		switch {
		case strings.Contains(input, "new"):
			words := strings.Split(input, " ")
			name := ""
			if len(words) > 1 {
				name = words[1]
			} else {
				name = "Untitled"
			}
			newNote(name)
		case input == "exit":
			fmt.Println("Shutting down...\nGoodbye!")
			break Infinite
		case input == "help":
			helpMenu()
		default:
			fmt.Println("Unknown command.")
		}
		input = ""
	}
}

func newNote(name string) {
	fmt.Println("Please enter the text for your note:")
	stdin := bufio.NewReader(os.Stdin)
	input, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimRight(input, " \n")
	saveNote(name, input)
}

func saveNote(name string, input string) {
	fmt.Printf("Saving note: '%s', with content: '%s'\n", name, input)
	fmt.Println("(Nothing happens yet.)")
}

func helpMenu() {
	fmt.Println("Welcome to the help menu!")
	showCommands()
}

func showCommands() {
	commands := []string{"new <note name>", "help", "exit"}
	fmt.Println("The available commands are:")
	for _, command := range commands {
		fmt.Println(command)
	}
}
