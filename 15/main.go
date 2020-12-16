package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"strconv"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for _, v := range tokens {
			val, _ := strconv.Atoi(v)
			numbers = append(numbers, val)
		}
	}
	fmt.Println("Answer 1:")
	fmt.Println(game(numbers, 2020))
	fmt.Println("Answer 2:")
	fmt.Println(game(numbers, 30000000))
}


func game(input []int, turns int) int {
	memory := make(map[int]int)//track mentions
	currentNumber := 0
	for i, v := range input {
		currentNumber = v
		memory[v] = i + 1 //turn starts at 1
	}
	for prev := len(input); prev < turns; prev = prev + 1 {
		currentLastSpokenTurn, mention := memory[currentNumber]
		memory[currentNumber] = prev
		if mention {
			currentNumber = prev - currentLastSpokenTurn
		} else {
			currentNumber = 0
		}
	}
	return currentNumber
}
