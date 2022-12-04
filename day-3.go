package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Link to Day 3 problem: https://adventofcode.com/2022/day/3
func main() {
	priorityMapping := make(map[string]int)
	priorityMapping["a"] = 1
	priorityMapping["b"] = 2
	priorityMapping["c"] = 3
	priorityMapping["d"] = 4
	priorityMapping["e"] = 5
	priorityMapping["f"] = 6
	priorityMapping["g"] = 7
	priorityMapping["h"] = 8
	priorityMapping["i"] = 9
	priorityMapping["j"] = 10
	priorityMapping["k"] = 11
	priorityMapping["l"] = 12
	priorityMapping["m"] = 13
	priorityMapping["n"] = 14
	priorityMapping["o"] = 15
	priorityMapping["p"] = 16
	priorityMapping["q"] = 17
	priorityMapping["r"] = 18
	priorityMapping["s"] = 19
	priorityMapping["t"] = 20
	priorityMapping["u"] = 21
	priorityMapping["v"] = 22
	priorityMapping["w"] = 23
	priorityMapping["x"] = 24
	priorityMapping["y"] = 25
	priorityMapping["z"] = 26
	priorityMapping["A"] = 27
	priorityMapping["B"] = 28
	priorityMapping["C"] = 29
	priorityMapping["D"] = 30
	priorityMapping["E"] = 31
	priorityMapping["F"] = 32
	priorityMapping["G"] = 33
	priorityMapping["H"] = 34
	priorityMapping["I"] = 35
	priorityMapping["J"] = 36
	priorityMapping["K"] = 37
	priorityMapping["L"] = 38
	priorityMapping["M"] = 39
	priorityMapping["N"] = 40
	priorityMapping["O"] = 41
	priorityMapping["P"] = 42
	priorityMapping["Q"] = 43
	priorityMapping["R"] = 44
	priorityMapping["S"] = 45
	priorityMapping["T"] = 46
	priorityMapping["U"] = 47
	priorityMapping["V"] = 48
	priorityMapping["W"] = 49
	priorityMapping["X"] = 50
	priorityMapping["Y"] = 51
	priorityMapping["Z"] = 52

	content, err := os.ReadFile("day-3-inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputsArray := strings.Split(string(content), "\n")
	inputsArray = inputsArray[:len(inputsArray)-1]
	fmt.Println(len(inputsArray))

	totalPoints := 0

	for i := 0; i < len(inputsArray); i = i + 3 {
		firstElfRuckSack := inputsArray[i]
		secondElfRuckSack := inputsArray[i+1]
		thirdElfRuckSack := inputsArray[i+2]

		firstElfCompartmentArray := strings.Split(firstElfRuckSack, "")
		secondElfCompartmentArray := strings.Split(secondElfRuckSack, "")
		thirdElfCompartmentArray := strings.Split(thirdElfRuckSack, "")

		commonLetter := ""
		var firstElfRuckSackSet = make(map[string]bool)
		var secondElfRuckSackSet = make(map[string]bool)
		for j := 0; j < len(firstElfCompartmentArray); j++ {
			firstElfRuckSackSet[firstElfCompartmentArray[j]] = true
		}
		for k := 0; k < len(secondElfCompartmentArray); k++ {
			secondElfRuckSackSet[secondElfCompartmentArray[k]] = true
		}
		for l := 0; l < len(thirdElfCompartmentArray); l++ {
			if firstElfRuckSackSet[thirdElfCompartmentArray[l]] == true && secondElfRuckSackSet[thirdElfCompartmentArray[l]] {
				commonLetter = thirdElfCompartmentArray[l]
				totalPoints = totalPoints + priorityMapping[commonLetter]
				break
			}
		}
	}
	fmt.Println(totalPoints)
}
