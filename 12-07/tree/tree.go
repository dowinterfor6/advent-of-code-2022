package tree

import "fmt"

type NodeFile struct {
	Name string
	Size int64
}

type Node struct {
	File    NodeFile
	DirName string
	Leaves  []*Node
	Parent  *Node
	// Type    string // how do I do "file" or "dir"?
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func FindNodeByDirName(dirName string, node *Node) *Node {
	var foundNode *Node
	for _, v := range node.Leaves {
		if v.DirName == dirName {
			foundNode = v
			break
		}
	}

	return foundNode
}

func CreateDirNode(dirName string, parent *Node) *Node {
	return &Node{DirName: dirName, Parent: parent}
}

func CreateFileNode(name string, size int64, parent *Node) *Node {
	return &Node{File: NodeFile{Name: name, Size: size}, Parent: parent}
}

func NodeToString(node *Node) {
	if node.DirName != "" {
		fmt.Printf("Dir Node  | DirName: %v | Num Leaves: %v", node.DirName, len(node.Leaves))
	} else {
		fmt.Printf("File Node | FileName: %v | FileSize: %v", node.File.Name, node.File.Size)
	}
	fmt.Println()
}
