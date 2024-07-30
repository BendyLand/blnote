package node

import (
	"fmt"
	"slices"
	"strings"
)

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
	var links []string
	var nonLinks []string
	for _, node := range nodes {
		var link *string
		var line string
		link = node.Link
		if link == nil {
			temp := ""
			link = &temp
		}
		if *link == "" && !slices.Contains(links, node.Name) {
			line = node.Name 
			nonLinks = append(nonLinks, line)
		} else {
			if !slices.Contains(links, node.Name) {
				line = *node.Link + " <- " + node.Name
				links = append(links, line)
			}
		}
	}
	nonLinks = removeLinksFromNonLinkGroup(links, nonLinks)
	tempSlice := slices.Concat(nonLinks, links)
	result := strings.Join(tempSlice, "\n")
	fmt.Println(result)
}

func removeLinksFromNonLinkGroup(links []string, nonLinks []string) []string {
	var result []string
	Outer: for _, item := range nonLinks {
		item = strings.TrimRight(item, "\n")
		for _, link := range links {
			link = strings.TrimRight(link, "\n")
			if strings.Contains(link, item) {
				continue Outer
			}
		}
		result = append(result, item)
	}
	return result
}

func GetNode(nodeName string, nodes Nodes) (*Node, error) {
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

func DeleteNode(nodeName string, nodes *Nodes) error {
	idx := -1
	for i := 0; i < len(*nodes); i++ {
		if (*nodes)[i].Name == nodeName {
			idx = i
			break
		}
	}
	if idx == -1 || idx >= len(*nodes) {
		err := fmt.Errorf("Node not found.\n")
		return err
	}
	*nodes = removeNode(*nodes, idx)
	fmt.Println("Node removed successfully.")
	return nil
}

func removeNode(nodes Nodes, i int) Nodes {
	nodes[i] = nodes[len(nodes)-1]
	return  nodes[:len(nodes)-1]
}

func CheckNode(name string, nodes Nodes) {
	for _, node := range nodes {
		if node.Name == name {
			fmt.Printf("%s: %s\n", node.Name, node.Text)
			return
		}
	}
	fmt.Println("Node not found.")
}

func EditNode(name string, newText string, nodes *Nodes) {
	for i := 0; i < len(*nodes); i++ {
		if (*nodes)[i].Name == name {
			(*nodes)[i].Text = newText
			fmt.Println("Node updated!")
			return
		}
	}
	fmt.Println("Node not found.")
}
