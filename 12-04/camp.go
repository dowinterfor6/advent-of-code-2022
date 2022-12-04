package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitOnChar(str string, char string) (string, string) {
	splitOnChar := strings.Split(str, char)
	return splitOnChar[0], splitOnChar[1]
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

	overlappingRanges := 0

	for scanner.Scan() {
		line := scanner.Text()
		leftAssignment, rightAssignment := splitOnChar(line, ",")
		left1, left2 := splitOnChar(leftAssignment, "-")
		right1, right2 := splitOnChar(rightAssignment, "-")

		l1, _ := strconv.ParseInt(left1, 10, 64)
		l2, _ := strconv.ParseInt(left2, 10, 64)
		r1, _ := strconv.ParseInt(right1, 10, 64)
		r2, _ := strconv.ParseInt(right2, 10, 64)

		leftRange := [2]int64{l1, l2}
		rightRange := [2]int64{r1, r2}

		if (leftRange[0] <= rightRange[0] && leftRange[1] >= rightRange[1]) || (rightRange[0] <= leftRange[0] && rightRange[1] >= leftRange[1]) {
			overlappingRanges++
		}
	}

	fmt.Printf("Complete Overlapping Ranges = %v", overlappingRanges)
	fmt.Println()
}

func part2() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	overlappingRanges := 0

	for scanner.Scan() {
		line := scanner.Text()
		leftAssignment, rightAssignment := splitOnChar(line, ",")
		left1, left2 := splitOnChar(leftAssignment, "-")
		right1, right2 := splitOnChar(rightAssignment, "-")

		l1, _ := strconv.ParseInt(left1, 10, 64)
		l2, _ := strconv.ParseInt(left2, 10, 64)
		r1, _ := strconv.ParseInt(right1, 10, 64)
		r2, _ := strconv.ParseInt(right2, 10, 64)

		leftRange := [2]int64{l1, l2}
		rightRange := [2]int64{r1, r2}

		// This feels like such a terrible solution but it works
		if (leftRange[0] <= rightRange[0] && leftRange[1] >= rightRange[0]) || (leftRange[0] <= rightRange[1] && leftRange[1] >= rightRange[1]) || (rightRange[0] <= leftRange[0] && rightRange[1] >= leftRange[0]) || (rightRange[0] <= leftRange[1] && rightRange[1] >= leftRange[1]) {
			overlappingRanges++
		}
	}

	fmt.Printf("Overlapping Ranges = %v", overlappingRanges)
	fmt.Println()
}
