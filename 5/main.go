package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var passes []string
	for scanner.Scan() {
		pass := scanner.Text()
		passes = append(passes, pass)
	}

	//problem 1
	var seats []int
	maxId := 0
	for _, pass := range passes {
		row:=0
		col:=0
		for i, r := range pass[0:7] {
			if r == 'F' { } else {
				row = row + int(math.Pow(float64(2), float64(6-i)))
			}
		}
		for i, c := range pass[7:10] {
			if c == 'L' { } else {
				col = col + int(math.Pow(float64(2), float64(2-i)))
			}
		}
		seatId := row * 8 + col
		if seatId > maxId { maxId = seatId }
		seats = append(seats, seatId)
	}
	fmt.Println("Max seat ID in dataset:")
	fmt.Println(maxId)
	sort.Ints(seats)
	for i, seat := range seats {
		if i < len(seats)-1 && seat+1 < seats[i+1] {
			fmt.Println(seat+1)
			fmt.Println(seats[i-1], seats[i], seats[i+1])
		}
	}
}
