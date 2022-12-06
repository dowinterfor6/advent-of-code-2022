package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
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
	hasDataStarted := false
	var inverseStackMap = make(map[int]*stack.Stack)
	var stackMap = make(map[int]*stack.Stack)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			hasDataStarted = true
			stackMap = *invertStackMap(&inverseStackMap)
			// Sanity check, since it's a map the keys are unsorted
			// printMap(&stackMap)
			// fmt.Println()
			continue
		}

		if !hasDataStarted {
			// Just setting up the stacks
			for i, v := range line {
				// Break when it reaches the number line
				if v == '1' {
					break
				}

				// The actual data lives in cols 1, 5, 9, ...
				if (i-1)%4 == 0 {
					idx := (i - 1) / 4

					if inverseStackMap[idx] == nil {
						inverseStackMap[idx] = stack.NewStack()
					}

					// Empty entries can exist in 'data' spots, this is to remove them
					if v != ' ' {
						inverseStackMap[idx].Push(string(v))
					}
				}
			}
		} else {
			// Actually moving stuff now
			strArr := strings.Split(line, " ")
			numToMove, err1 := strconv.ParseInt(strArr[1], 10, 64)
			check(err1)
			fromStackNum, err2 := strconv.ParseInt(strArr[3], 10, 64)
			check(err2)
			toStackNum, err3 := strconv.ParseInt(strArr[5], 10, 64)
			check(err3)

			for i := 0; i < int(numToMove); i++ {
				currStack := stackMap[int(fromStackNum-1)]
				stackMap[int(toStackNum-1)].Push(currStack.Peek().Value)
				currStack.Pop()
			}
		}
	}

	printTopMap(&stackMap)
}

// TBH I couldn't figure out how to invert a stack map from src to dest as
// param pointers, so I just made a new one and returned it from the function
func invertStackMap(src *map[int]*stack.Stack) *map[int]*stack.Stack {
	var invertedStackMap = make(map[int]*stack.Stack)

	for key, v := range *src {
		for v.Length() > 0 {
			if invertedStackMap[key] == nil {
				invertedStackMap[key] = stack.NewStack()
			}
			invertedStackMap[key].Push(v.Peek().Value)
			v.Pop()
		}
	}

	return &invertedStackMap
}

func printMap(mapToPrint *map[int]*stack.Stack) {
	for k, v := range *mapToPrint {
		fmt.Printf("%v: %v", k, v.ToString())
		fmt.Println()
	}
}

// I tried to use a for loop but couldn't get it to work, typing issues
func printTopMap(mapToPrint *map[int]*stack.Stack) {
	// Also this isn't dynamic, this approach probably won't ever be dynamic
	slice := make([]string, 9)

	for k, v := range *mapToPrint {
		slice[k] = v.Top.Value
	}

	fmt.Println(strings.Join(slice, ""))
}

func part2() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hasDataStarted := false
	var inverseStackMap = make(map[int]*stack.Stack)
	var stackMap = make(map[int]*stack.Stack)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			hasDataStarted = true
			stackMap = *invertStackMap(&inverseStackMap)
			// Sanity check, since it's a map the keys are unsorted
			// printMap(&stackMap)
			fmt.Println()
			continue
		}

		if !hasDataStarted {
			// Just setting up the stacks
			for i, v := range line {
				// Break when it reaches the number line
				if v == '1' {
					break
				}

				// The actual data lives in cols 1, 5, 9, ...
				if (i-1)%4 == 0 {
					idx := (i - 1) / 4

					if inverseStackMap[idx] == nil {
						inverseStackMap[idx] = stack.NewStack()
					}

					// Empty entries can exist in 'data' spots, this is to remove them
					if v != ' ' {
						inverseStackMap[idx].Push(string(v))
					}
				}
			}
		} else {
			// Actually moving stuff now
			strArr := strings.Split(line, " ")
			numToMove, err1 := strconv.ParseInt(strArr[1], 10, 64)
			check(err1)
			fromStackNum, err2 := strconv.ParseInt(strArr[3], 10, 64)
			check(err2)
			toStackNum, err3 := strconv.ParseInt(strArr[5], 10, 64)
			check(err3)

			movingStack := stack.NewStack()

			for i := 0; i < int(numToMove); i++ {
				currStack := stackMap[int(fromStackNum-1)]
				movingStack.Push(currStack.Peek().Value)
				currStack.Pop()
			}

			movesNeeded := movingStack.Length()
			for i := 0; i < movesNeeded; i++ {
				stackMap[int(toStackNum-1)].Push(movingStack.Peek().Value)
				movingStack.Pop()
			}
		}
	}

	printTopMap(&stackMap)
}
