package main

import (
	"bufio"
	"fmt"
	"os"
)

// Reads in input line by line and places it in a 2D array of strings
func read_file(name string) [][]string {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)

	defer file.Close()

	var items [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var temp_array []string
		for _, item := range line {
			temp_array = append(temp_array, string(item))
		}
		items = append(items, temp_array)
	}
	return items
}

func main() {
	// First the input is read
	lines := read_file("input")

	// Then it sets variables to see how many rolls were accessible total or how many are currently accessible
	accessable := 0
	curr_accessable := -1

	// It then iterates through the problem until there are no rolls that are accessible
	for curr_accessable != 0 {
		curr_accessable = 0

		// It iterates through every roll, checking if it is a valid roll that you can check
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[y]); x++ {
				rolls := 0
				if lines[y][x] == "@" {

					// Then it iterates through all of the adjacent rolls
					for deltay := -1; deltay < 2; deltay++ {
						for deltax := -1; deltax < 2; deltax++ {

							// It ignores if the delta is (0,0) because then it would be the current roll
							if !(deltax == deltay && deltax == 0) {

								// If the position is not (0,0) then it checks to make sure that the (x,y) pair created
								// is within the bounds of the 2D array so as to not cause exceptions
								if (x+deltax) >= 0 && (x+deltax) <= (len(lines[y])-1) && (y+deltay) >= 0 && (y+deltay) <= (len(lines)-1) {

									// If it is within bounds, then it checks to make sure its a roll "@" or
									// a roll about to be removed "x" and then increments the rolls value which
									// represents the number of adjacent rolls
									if lines[y+deltay][x+deltax] == "@" || lines[y+deltay][x+deltax] == "x" {
										rolls++
									}
								}
							}
						}
					}

					// If the number of adjacent rolls is under 4, then it is added to the currently accessible rolls and
					// is prepped to be removed
					if rolls < 4 {
						curr_accessable++
						lines[y][x] = "x"
					}
				}
			}
		}

		// Once all of the rolls have been checked, it then removes all of the rolls, replacing them with blank spaces
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[y]); x++ {
				if lines[y][x] == "x" {
					lines[y][x] = "."
				}
			}
		}

		// The number of accessible rolls is then added and printed
		accessable += curr_accessable
	}
	fmt.Println(accessable)
}
