package note

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetText() string {
	fmt.Println("Please enter the text for your note:")
	stdin := bufio.NewReader(os.Stdin)
	input, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	input = strings.TrimRight(input, " \n")
	return input
}

func HelpMenu() {
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
