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
	/* 
	todo:
	Given nodes: two <- one, if node `two` is deleted, when the nodes are displayed, node `one` will still show its Link field properly, so it will appear like the parent itself still exists.
	If you attempt to `check` the contents of the parent node, it will not be found. 
	todo: find way to validate Link before displaying, possibly display updated message when removed or remove all links when parent is removed.
	*/
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
			newNode := createNewNode(input)
			node.UpdateNodes(newNode, &nodes)
		case strings.Contains(input, "remove"):
			removeNode(input, err, &nodes)
		case strings.Contains(input, "edit"):
			args := strings.Split(input, " ")
			if len(args) < 2 {
				fmt.Println("Unable to edit node. No name given.")
				continue
			}
			fmt.Println("Enter new text for node:")
			newText, err := stdin.ReadString('\n')
			if err != nil {
				fmt.Println("Error geting new text for node.")
				continue
			}
			newText = strings.TrimRight(newText, " \n")
			node.EditNode(args[1], newText, &nodes)
		case input == "help":
			note.HelpMenu()
		case input == "show":
			node.DisplayNodes(nodes)
		case strings.Contains(input, "link"):
			linkNodes(input, nodes)
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

func createNewNode(input string) node.Node {
	words := strings.Split(input, " ")
	name := ""
	if len(words) > 1 {
		name = words[1]
	} else {
		name = "Untitled"
	}
	text := note.GetText()
	return node.NewNode(name, text, nil)
}

func removeNode(input string, err error, nodes *node.Nodes) {
	args := strings.Split(input, " ")
	if len(args) > 1 {
		args = args[1:]
		for _, arg := range args {
			err = node.DeleteNode(arg, nodes)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Unable to delete node. Invalid arguments.")
	}
}

func linkNodes(input string, nodes node.Nodes) {
	inputNodes := strings.Split(input, " ")
	if len(inputNodes) > 2 {
		node1, err1 := node.GetNode(inputNodes[1], nodes)
		node2, err2 := node.GetNode(inputNodes[2], nodes)
		if err1 != nil || err2 != nil {
			fmt.Println("Error getting first two nodes:", err1, err2, "| Aborting link.")
			return
		}
		node.LinkNodes(node1, node2)
		fmt.Println("Nodes linked successfully!")
	} else {
		fmt.Println("Unable to link nodes. Invalid arguments.")
	}
}
