package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./data.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mostCalories := 0
	currTotalCalories := 0

	var insertionSlice []int
	var sortSlice []int

	for scanner.Scan() {
		// strconv attempts to convert string to integer
		// data file lines contain either stringified integers
		// or new line \n, which will cause an error, which
		// also indicates the end of a chunk of data
		i, err := strconv.Atoi(scanner.Text());

		if (err != nil) {
			if (currTotalCalories > mostCalories) {
				mostCalories = currTotalCalories
			}
			
			insertionSlice = insertionSort(insertionSlice, currTotalCalories);
			sortSlice = append(sortSlice, currTotalCalories)
			currTotalCalories = 0
		} else {
			currTotalCalories += i
		}
	}
		
	fmt.Printf("Most Calories: %v \n", mostCalories)

	// Using built-in sort (210957)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i] > sortSlice[j]
	})

	topThreeCaloriesSort := 0
	for _, v := range sortSlice[:3] {
		topThreeCaloriesSort += v
	}
	fmt.Printf("Top 3 Calories (sort): %v \n", topThreeCaloriesSort)
	
	// Using insertion sort (210957)
	topThreeCalories := 0
	for _, v := range insertionSlice[:3] {
		topThreeCalories += v
	}
	fmt.Printf("Top 3 Calories: %v \n", topThreeCalories)
}

func insertionSort(slice []int, value int) []int {
	// append and copy will react differently depending on size/capacity
	// to be on the safe side, I'll be making copies with correct size and
	// capacities
	sliceCopy := make([]int, len(slice))
	if (len(slice) > 0) {
		_ = copy(sliceCopy, slice)
	}

	if (len(sliceCopy) == 0) {
		return append(sliceCopy, value)
	}

	for i, v := range sliceCopy {
		if (value > v) {
			leftSlice := make([]int, i)
			_ = copy(leftSlice, sliceCopy[:i])
			leftSlice = append(leftSlice, value)

			allSlice := make([]int, len(sliceCopy[i:]), len(sliceCopy) + 1)
			_ = copy(allSlice, sliceCopy[i:])
			allSlice = append(leftSlice, allSlice...)
			
			return allSlice
		}
	}

	return append(sliceCopy, value)
}