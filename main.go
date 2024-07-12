package main

import (
	"blnote/node"
	"blnote/note"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nodes := node.Nodes{}
	fmt.Println("Welcome to blnote!")
	mainLoop(nodes)
}

func mainLoop(nodes node.Nodes) {
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
			text := note.GetText()
			newNode := node.NewNode(name, text)
			node.UpdateNodes(newNode, &nodes)
		case input == "exit":
			fmt.Println("Shutting down...\nGoodbye!")
			break Infinite
		case input == "help":
			note.HelpMenu()
		case input == "show":
			node.DisplayNodes(nodes)
		default:
			fmt.Println("Unknown command.")
		}
	}
}
