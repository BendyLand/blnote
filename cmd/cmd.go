package cmd

import (
	"blnote/node"
)

func SaveNote(name string, text string) node.Node {
	return node.Node{Name: name, Text: text}
}
