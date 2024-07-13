package node

import "fmt"

type Node struct {
	Name string
	Text string
	Link *string 
}

type Nodes []Node

func NewNode(name string, text string, link *string) Node {
	return Node{name, text, link}
}

func UpdateNodes(node Node, nodes *Nodes) {
	*nodes = append(*nodes, node)
}

func DisplayNodes(nodes Nodes) {
	fmt.Println("Current nodes:")
	for _, node := range nodes {
		var link *string
		link = node.Link
		if link == nil {
			temp := "nil"
			link = &temp
		}
		fmt.Printf("Name: %s\nLink: %s\nText: %s\n\n", node.Name, *link, node.Text)
	}
}

func GetNode(nodeName string, nodes Nodes) (*Node, error){
	// using a range loop breaks the behavior.
	for i := 0; i < len(nodes); i++ { 
		if nodes[i].Name == nodeName {
			return &nodes[i], nil
		}
	}
	err := fmt.Errorf("Node not found.")
	return nil, err
}

func LinkNodes(node1 *Node, node2 *Node) {
	if node2.Link == nil {
		node2.Link = new(string)
	}
	*node2.Link = node1.Name
}
