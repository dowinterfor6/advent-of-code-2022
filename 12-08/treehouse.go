package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
}

// Alternative approach
// Make a data structure that keeps track of max per col and row
func part1() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputArr := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			inputArr = append(inputArr, line)
		}
	}

	numVisiblePositions := checkVisiblePositions(inputArr)

	fmt.Printf("# visible treehouse positions: %v", numVisiblePositions)
	fmt.Println()
}

func checkVisiblePositions(arr []string) int {
	visiblePositions := 0

	for rowIdx, row := range arr {
		if rowIdx == 0 || rowIdx == len(arr)-1 {
			fmt.Println("Top or Bottom row")
			visiblePositions += len(arr)
			continue
		}

		for colIdx, v := range row {
			if colIdx == 0 || colIdx == len(row)-1 {
				fmt.Println()
				fmt.Println("Left or Right col")
				visiblePositions++
				continue
			}

			currVal, err := strconv.ParseInt(string(v), 10, 64)
			check(err)
			fmt.Println()
			fmt.Printf("currVal: %v | ", currVal)

			isVisibleFromTop := true
			topIdx := rowIdx - 1
			for topIdx >= 0 {
				topVal, err := strconv.ParseInt(string(arr[topIdx][colIdx]), 10, 64)
				fmt.Print(topVal)
				check(err)

				if topVal >= currVal {
					isVisibleFromTop = false
					break
				}
				topIdx--
			}
			fmt.Print(" | ")

			if isVisibleFromTop {
				fmt.Print("Top Valid, adding")
				visiblePositions++
				continue
			}

			isVisibleFromBot := true
			botIdx := rowIdx + 1
			for botIdx < len(arr) {
				botVal, err := strconv.ParseInt(string(arr[botIdx][colIdx]), 10, 64)
				fmt.Print(botVal)
				check(err)

				if botVal >= currVal {
					isVisibleFromBot = false
					break
				}
				botIdx++
			}
			fmt.Print(" | ")

			if isVisibleFromBot {
				fmt.Print("Bot Valid, adding")
				visiblePositions++
				continue
			}

			isVisibleFromRight := true
			rightIdx := colIdx + 1
			for rightIdx < len(row) {
				rightVal, err := strconv.ParseInt(string(arr[rowIdx][rightIdx]), 10, 64)
				fmt.Print(rightVal)
				check(err)

				if rightVal >= currVal {
					isVisibleFromRight = false
					break
				}
				rightIdx++
			}
			fmt.Print(" | ")

			if isVisibleFromRight {
				fmt.Print("Right Valid, adding")
				visiblePositions++
				continue
			}

			isVisibleFromLeft := true
			leftIdx := colIdx - 1
			for leftIdx >= 0 {
				leftVal, err := strconv.ParseInt(string(arr[rowIdx][leftIdx]), 10, 64)
				fmt.Print(leftVal)
				check(err)

				if leftVal >= currVal {
					isVisibleFromLeft = false
					break
				}
				leftIdx++
			}
			fmt.Print(" | ")

			if isVisibleFromLeft {
				fmt.Print("Left Valid, adding")
				visiblePositions++
				continue
			}
		}
	}

	return visiblePositions
}
