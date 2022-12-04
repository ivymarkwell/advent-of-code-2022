package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Link to Day 4 problem: https://adventofcode.com/2022/day/4
func main() {
	content, err := os.ReadFile("day-4-inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputsArray := strings.Split(string(content), "\n")
	inputsArray = inputsArray[:len(inputsArray)-1]

	overlappingNumPairs := 0
	for i := 0; i < len(inputsArray); i++ {
		elfAssignments := strings.Split(inputsArray[i], ",")

		elf1Assignment := elfAssignments[0]
		elf1AssignmentSections := strings.Split(elf1Assignment, "-")

		elf2Assignment := elfAssignments[1]
		elf2AssignmentSections := strings.Split(elf2Assignment, "-")

		elf1AssignmentSectionStart, _ := strconv.Atoi(elf1AssignmentSections[0])
		elf1AssignmentSectionEnd, _ := strconv.Atoi(elf1AssignmentSections[1])
		elf2AssignmentSectionStart, _ := strconv.Atoi(elf2AssignmentSections[0])
		elf2AssignmentSectionEnd, _ := strconv.Atoi(elf2AssignmentSections[1])

		if elf1AssignmentSectionStart <= elf2AssignmentSectionStart && elf1AssignmentSectionEnd >= elf2AssignmentSectionEnd {
			overlappingNumPairs++
		} else if elf1AssignmentSectionStart >= elf2AssignmentSectionStart && elf1AssignmentSectionEnd <= elf2AssignmentSectionEnd {
			overlappingNumPairs++
		} else if elf1AssignmentSectionStart <= elf2AssignmentSectionStart && elf1AssignmentSectionEnd >= elf2AssignmentSectionStart {
			overlappingNumPairs++
		} else if elf1AssignmentSectionStart >= elf2AssignmentSectionStart && elf2AssignmentSectionEnd >= elf1AssignmentSectionStart {
			overlappingNumPairs++
		} else if elf1AssignmentSectionEnd == elf2AssignmentSectionStart {
			overlappingNumPairs++
		} else if elf1AssignmentSectionStart == elf2AssignmentSectionEnd {
			overlappingNumPairs++
		} else {
			overlappingNumPairs = overlappingNumPairs
		}
	}
	fmt.Println(overlappingNumPairs)
}
