package node

import "fmt"

type Node struct {
	Name string
	Text string
}

type Nodes []Node

func NewNode(name string, text string) Node {
	return Node{name, text}
}

func UpdateNodes(node Node, nodes *Nodes) {
	*nodes = append(*nodes, node)
}

func DisplayNodes(nodes Nodes) {
	fmt.Println("Current nodes:")
	for _, node := range nodes {
		fmt.Printf("%s: %s\n", node.Name, node.Text)
	}
}
