package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tree"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
}

func part1() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	directoryTree := tree.NewTree()
	currNode := tree.CreateDirNode("/", nil)
	directoryTree.Root = currNode
	scanner.Scan()
	line := scanner.Text()

	for line != "" {
		// Read commands (anything that starts with $)
		fmt.Println()
		fmt.Printf("Input: %v", line)
		fmt.Println()

		if strings.Contains(line, "$ cd") {
			if strings.Contains(line, "/") {
				scanner.Scan()
				line = scanner.Text()
				currNode = directoryTree.Root
				continue
			} else if strings.Contains(line, "..") {
				scanner.Scan()
				line = scanner.Text()
				currNode = currNode.Parent
				continue
			} else {
				foundNode := tree.FindNodeByDirName(strings.Split(line, "$ cd ")[1], currNode)
				if foundNode == nil {
					panic("Could not find node by dir name")
				}
				currNode = foundNode
				scanner.Scan()
				line = scanner.Text()
			}
		} else if strings.Contains(line, "$ ls") {
			scanner.Scan()
			line = scanner.Text()

			for !strings.Contains(line, "$") && line != "" {
				fmt.Println(line)
				fmt.Println()
				if currNode.Leaves == nil {
					currNode.Leaves = make([]*tree.Node, 0)
				}

				var newNode *tree.Node
				if strings.Contains(line, "dir ") {
					newNode = tree.CreateDirNode(strings.Split(line, "dir ")[1], currNode)
				} else {
					fileInputs := strings.Split(line, " ")
					fileSize, err := strconv.ParseInt(fileInputs[0], 10, 32)
					check(err)
					newNode = tree.CreateFileNode(fileInputs[1], fileSize, currNode)
				}
				currNode.Leaves = append(currNode.Leaves, newNode)

				scanner.Scan()
				line = scanner.Text()
			}
		} else {
			fmt.Println("How the hell did you get here")
			panic("RIP")
		}
	}
}
