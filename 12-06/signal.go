package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
	part2()
}

// This solution is very hard coded to question prompt
func part1() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := range line {
			if i == len(line)-5 {
				fmt.Println("Reached end of input without marker")
				break
			}

			substr := line[i : i+4]
			invalid := false

			for i, testChar := range substr {
				var testSubstr string
				switch i {
				case 0:
					testSubstr = substr[1:]
				case 3:
					testSubstr = substr[:3]
				default:
					testSubstr = substr[:i] + substr[i+1:]
				}

				if strings.Contains(testSubstr, string(testChar)) {
					invalid = true
				}
			}

			if !invalid {
				fmt.Printf("First packet marker ends after char: %v", i+4)
				fmt.Println()
				break
			}
		}
	}
}

func part2() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := range line {
			if i == len(line)-15 {
				fmt.Println("Reached end of input without marker")
				break
			}

			substr := line[i : i+14]
			invalid := false

			for i, testChar := range substr {
				var testSubstr string
				switch i {
				case 0:
					testSubstr = substr[1:]
				case 13:
					testSubstr = substr[:13]
				default:
					testSubstr = substr[:i] + substr[i+1:]
				}

				if strings.Contains(testSubstr, string(testChar)) {
					invalid = true
				}
			}

			if !invalid {
				fmt.Printf("First message marker ends after char: %v", i+14)
				fmt.Println()
				break
			}
		}
	}
}
