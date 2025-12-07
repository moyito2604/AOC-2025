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

func main() {
    
    // First, the file is read in
    rotations := read_file("input")
    
    //Variables are defined. at_zero is the amount of times it has crossed zero
    at_zero := 0
    currentpos := 50
    startpos := 0
    endingpos := 99

    // Rotations are then performed based on the input given
    for _, rotation := range rotations {

        // The rotation is split into direction and amount of rotations
        direction := rotation[0:1]
        amount, _ := strconv.Atoi(rotation[1:])
        switch direction {
        case "L":

            // Then it iterates through the amount of rotations, seeing applying the changes one tick at a time
            for counter := 0; counter < amount; counter++{
                currentpos -= 1
                if currentpos < startpos {
                    currentpos = endingpos
                }
                // Solution for Part One Requires commenting this
                if currentpos == 0 {
                    at_zero += 1
                }
            }
        case "R":

            // It also applies changes one tick at a time for right handed rotations as well
            for counter := 0; counter < amount; counter++{
                currentpos += 1
                if currentpos > endingpos {
                    currentpos = startpos
                }
                // Solution for Part One Requires commenting this
                if currentpos == 0 {
                    at_zero += 1
                }
            }
        }
        // Solution for Part One requires this
        // if currentpos == 0 {
        //     at_zero += 1
        // }

        // The current position of the lock is then printed
        fmt.Println(currentpos)
    }

    // The solution is then printed at the end
    fmt.Println(at_zero)
}