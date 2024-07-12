package cmd

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"blnote/note"
	"blnote/node"
)


func SaveNote(name string, text string) node.Node {
	return node.Node{name, text}
}
