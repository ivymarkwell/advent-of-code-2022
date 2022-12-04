package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Link to Day 1 problem: https://adventofcode.com/2022/day/1
func main() {

	content, err := os.ReadFile("day-1-inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputsArray := strings.Split(string(content), "\n")

	// current total calories an elf is carrying
	totalCaloriesCarried := 0

	// a list of the total number of calories each elf is carrying
	var caloriesArray []int
	for i := 0; i < len(inputsArray); i++ {
		// every elf is separated by an empty line
		// reset the totalCalories Carried for new elves and append their totalCaloriesCaried to the caloriesArray
		if inputsArray[i] == "" {
			caloriesArray = append(caloriesArray, totalCaloriesCarried)
			totalCaloriesCarried = 0
		} else {
			calories, err := strconv.Atoi(inputsArray[i])
			if err != nil {
				fmt.Println(err)
			} else {
				totalCaloriesCarried = totalCaloriesCarried + calories
			}
		}
	}
	sort.Ints(caloriesArray)

	maxCaloriesCarriedA := caloriesArray[len(caloriesArray)-1]
	maxCaloriesCarriedB := caloriesArray[len(caloriesArray)-2]
	maxCaloriesCarriedC := caloriesArray[len(caloriesArray)-3]

	fmt.Println(maxCaloriesCarriedA + maxCaloriesCarriedB + maxCaloriesCarriedC)
}
