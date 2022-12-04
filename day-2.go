package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Link to Day 2 problem: https://adventofcode.com/2022/day/2
func main() {

	content, err := os.ReadFile("day-2-inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputsArray := strings.Split(string(content), "\n")
	inputsArray = inputsArray[:len(inputsArray)-1]

	totalPoints := 0

	// rock A
	// paper B
	// scissors C
	for i := 0; i < len(inputsArray); i++ {
		round := inputsArray[i]
		roundHands := strings.Split(round, " ")

		opponentHand := roundHands[0]
		handToPlay := roundHands[1]

		if handToPlay == "X" && opponentHand == "C" {
			// we need to lose, we throw paper
			totalPoints = totalPoints + 2
		} else if handToPlay == "Y" && opponentHand == "B" {
			// we need to draw, we throw paper
			totalPoints = totalPoints + 2
		} else if handToPlay == "Z" && opponentHand == "A" {
			// we need to win, we throw paper
			totalPoints = totalPoints + 2
		} else if handToPlay == "X" && opponentHand == "A" {
			// we need to lose, we throw scissors
			totalPoints = totalPoints + 3
		} else if handToPlay == "Y" && opponentHand == "C" {
			// we need to draw, we throw scissors
			totalPoints = totalPoints + 3
		} else if handToPlay == "Z" && opponentHand == "B" {
			// we need to win, we throw scissors
			totalPoints = totalPoints + 3
		} else {
			totalPoints = totalPoints + 1
		}

		if handToPlay == "Z" {
			// winning round points
			totalPoints = totalPoints + 6
		} else if handToPlay == "Y" {
			// draw round points
			totalPoints = totalPoints + 3
		} else {
			// losing round points
			totalPoints = totalPoints
		}
	}
	fmt.Println(totalPoints)
}
