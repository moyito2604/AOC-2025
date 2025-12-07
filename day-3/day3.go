package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Reads in input line by line and places it in an array of strings
func read_file(name string) []string {
    file, _ := os.Open(name)
    scanner := bufio.NewScanner(file)

    defer file.Close()

    var items []string
    for scanner.Scan() {
        line := scanner.Text()
        items = append(items, line)
    }

    return items
}

// Finds the max item in a array of ints
func find_max(items []int) int{
	max := 0
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

func main() {
	
	// First the input is read
    lines := read_file("input")

	// The Joltage variable is defined and the amount of digits required is defined as well
	joltage := 0
	digits := 12

	// Then it iterates through the lines of batteries
	for _, line := range lines {
		start := 0
		var batteries []int
		var solution []int
		var solution_str string

		// It then converts all of the characters into an integer array for it to be easier to work with
		for _, item := range line {
			battery, _ := strconv.Atoi(string(item))
			batteries = append(batteries, battery)
		}

		// It then iterates through all of the digits required in the solution
		for digit := 1; digit <= digits; digit++ {

			// It sets the new end bound leaving enough for the rest of the digits
			end := len(batteries) - (digits - digit)

			// It then finds the max value with in the bounds of start and end and adds it to the solution
			max_value := find_max(batteries[start:end])
			solution = append(solution, max_value)

			// It then reassigns the start value to the position after it found the solution
			for counter := start; counter < len(batteries); counter++ {
				if batteries[counter] == max_value {
					start = counter + 1
					break
				}
			}
		}

		// Then, the solution is combined into a string
		for _, battery := range solution {
			solution_str += strconv.Itoa(battery)
		}

		// The solution is printed and then added to the total joltage value
		fmt.Println(solution_str)
		curr_joltage, _ := strconv.Atoi(solution_str)
		joltage += curr_joltage
	}

	// The total joltage value is then printed as well
	fmt.Println(joltage)
}