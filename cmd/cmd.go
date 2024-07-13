package cmd

import (
	"blnote/node"
	"blnote/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func SaveNote(name string, text string) node.Node {
	return node.Node{Name: name, Text: text}
}

func WriteNodesToFile(nodes node.Nodes) {
	exists, _ := utils.Exists("./NodeStorage")
	if !exists {
		err := os.MkdirAll("./NodeStorage", 0755)
		if err != nil {
			fmt.Println("Unable to create dir: './NodeStorage'")
			os.Exit(1)
		}
	}
	file, err := os.Create("./NodeStorage/nodes.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()
	errs := 0 
	file.WriteString("{\n")
	for i, node := range nodes {
		line := ""
		if node.Link != nil {
			line = fmt.Sprintf("\t\"%s.%s\": \"%s\"", *node.Link, node.Name, node.Text)
			if i < len(nodes)-1 {
				line += ",\n"
			} else {
				line += "\n"
			}
		} else {
			line = fmt.Sprintf("\t\"%s\": \"%s\"", node.Name, node.Text)
			if i < len(nodes)-1 {
				line += ",\n"
			} else {
				line += "\n"
			}
		}
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing string to file:", err)
			errs++
		} 
	}
	file.WriteString("}")
	if errs == 0 {
		fmt.Println("Nodes written successfully!")
	} else {
		fmt.Println("Nodes encountered problems being written. Please double check the contents.")
	}
}

func ReadNodesFromFile() node.Nodes {
	file, err := os.ReadFile("./NodeStorage/nodes.json")
	if err != nil {
		fmt.Println("No previous nodes found.")
		return node.Nodes{}
	}
	result := make(map[string]string)
	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println("Unable to parse JSON.")
		return node.Nodes{}
	}
	nodes := new(node.Nodes)
	for k, v := range result {
		name := ""
		link := ""
		if strings.Contains(k, ".") {
			parts := strings.Split(k, ".")
			name = parts[0]
			link = parts[1]
		} else {
			name = k
		}
		temp := node.Node{Name: name, Text: v, Link: &link}
		*nodes = append(*nodes, temp)
	}
	return *nodes
}
