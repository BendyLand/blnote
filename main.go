package main

import (
	"blnote/cmd"
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
			newNode := node.NewNode(name, text, nil)
			node.UpdateNodes(newNode, &nodes)
		case strings.Contains(input, "remove"):
			args := strings.Split(input, " ")
			if len(args) > 1 {
				args = args[1:]
				for _, arg := range args {
					err = node.DeleteNode(arg, &nodes)
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				fmt.Println("Unable to delete node. Invalid arguments.")
			}
		case input == "help":
			note.HelpMenu()
		case input == "show":
			node.DisplayNodes(nodes)
		case strings.Contains(input, "link"):
			inputNodes := strings.Split(input, " ")
			if len(inputNodes) > 2 {
				node1, err1 := node.GetNode(inputNodes[1], nodes)
				node2, err2 := node.GetNode(inputNodes[2], nodes)
				if err1 != nil || err2 != nil {
					fmt.Println("Error getting first two nodes:", err1, err2, "| Aborting link.")
					continue
				}
				node.LinkNodes(node1, node2)
				fmt.Println("Nodes linked successfully!")
			} else {
				fmt.Println("Unable to link nodes. Invalid arguments.")
			}
		case input == "read":
			nodes = cmd.ReadNodesFromFile()
		case strings.Contains(input, "check"): 
			args := strings.Split(input, " ")
			if len(args) < 2 {
				fmt.Println("Unable to check node. No name given.")
				continue
			} 
			node.CheckNode(args[1], nodes)
		case input == "exit":
			fmt.Println("Saving current nodes...")
			cmd.WriteNodesToFile(nodes)
			fmt.Println("Shutting down...\nGoodbye!")
			break Infinite
		default:
			fmt.Println("Unknown command.")
		}
	}
}
