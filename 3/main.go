package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var maze []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}

	treesEncountered := 0
	xIndex := 0 //start
	//Question one
	for i, s := range maze {
		if i == 0 {} else {
		xIndex = (xIndex + 3) % 31
		if s[xIndex] == '#' {
			treesEncountered = treesEncountered + 1
		}}
	}
	fmt.Println("Part 1 trees encountered:")
	fmt.Println(treesEncountered)
	answer := 1
	slopes := [][]int{{1,1},{3,1},{5,1},{7,1},{1,2},}
	for _, slope := range slopes {
		xIndex = 0
		yIndex := 0
		treesEncountered = 0
		for ok := true; ok; ok = yIndex<len(maze) {
			if yIndex == 0 {} else {
				xIndex = (xIndex + slope[0]) % 31
				if maze[yIndex][xIndex] == '#' {
					treesEncountered = treesEncountered + 1
				}
			}
			yIndex = yIndex + slope[1]
		}
		answer = answer * treesEncountered
	}

	fmt.Println("Part 2 answer:")
	fmt.Println(answer)
}
