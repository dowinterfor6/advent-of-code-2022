package main

import (
	"bufio"
	"fmt"
	"os"
	"set"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPriorityOfLetter(char string) int {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, v := range letters {
		if string(v) == char {
			return i + 1
		}
	}

	panic("Invalid char")
}

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prioritySum := 0

	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		left := line[:length/2]
		right := line[length/2:]

		leftSet := set.NewSet()

		for _, v := range left {
			leftSet.Add(string(v))
		}

		for _, v := range right {
			if leftSet.Contains(string(v)) {
				prioritySum += getPriorityOfLetter(string(v))
				break
			}
		}
	}

	fmt.Printf("Priority Sum: %v", prioritySum)
	fmt.Println()
}

func part2() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prioritySum := 0

	for scanner.Scan() {
		set1 := set.NewSet()
		set2 := set.NewSet()

		// This solution isn't dynamic, it only works for this specific number of lines
		// Prior attempt was at using a map, which failed since duplicate letters can
		// exist per line
		for i := 0; i < 3; i++ {

		lineloop:
			for _, v := range scanner.Text() {
				switch i {
				case 0:
					set1.Add(string(v))
				case 1:
					set2.Add(string(v))
				case 2:
					if set1.Contains(string(v)) && set2.Contains(string(v)) {
						prioritySum += getPriorityOfLetter(string(v))
						// This probably isn't "best practice" I think
						break lineloop
					}
				}
			}
			if i < 2 {
				scanner.Scan()
			}
		}

		fmt.Println()
	}

	fmt.Printf("Priority Sum 2: %v", prioritySum)
	fmt.Println()
}
