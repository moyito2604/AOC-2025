package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
    rotations := read_file("input")
    at_zero := 0
    currentpos := 50
    startpos := 0
    endingpos := 99
    for _, rotation := range rotations {
        direction := rotation[0:1]
        amount, _ := strconv.Atoi(rotation[1:])
        switch direction {
        case "L":
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
        fmt.Println(currentpos)
    }

    fmt.Println(at_zero)
}