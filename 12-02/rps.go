package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	A = Rock
	B = Paper
	C = Scissors

	X = Rock           1
	Y = Paper          2
	Z = Scissors       3

	Lose = 0
	Draw = 3
	Win = 6

	A X = Draw
	A Y = Win
	A Z = Lose
	B X = Lose
	B Y = Draw
	B Z = Win
	C X = Win
	C Y = Lose
	C Z = Draw
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	// This could've been done dynamically, or even required for a larger data set, but oh well
	scoreMap := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		i := scanner.Text()

		score += scoreMap[i]
	}

	fmt.Printf("P1 Total Score: %v", score)
	fmt.Println()
}

func part2() {
	ownChoiceScoreMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	outcomeScoreMap := map[string]int{
		"WIN":  6,
		"DRAW": 3,
		"LOSE": 0,
	}

	outcomeMap := map[string]string{
		"X": "LOSE",
		"Y": "DRAW",
		"Z": "WIN",
	}

	ownWinChoices := map[string]string{
		"A": "b",
		"B": "c",
		"C": "a",
	}

	ownLoseChoices := map[string]string{
		"A": "c",
		"B": "a",
		"C": "b",
	}

	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		i := scanner.Text()

		strArr := strings.Split(i, " ")
		p1Choice := strArr[0]
		p2Outcome := strArr[1]

		switch outcomeMap[p2Outcome] {
		case "LOSE":
			score += ownChoiceScoreMap[ownLoseChoices[p1Choice]] + outcomeScoreMap["LOSE"]
		case "DRAW":
			score += ownChoiceScoreMap[strings.ToLower(p1Choice)] + outcomeScoreMap["DRAW"]
		case "WIN":
			score += ownChoiceScoreMap[ownWinChoices[p1Choice]] + outcomeScoreMap["WIN"]
		}
	}

	fmt.Printf("P2 Total Score: %v", score)
	fmt.Println()
}
