package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	//part 1, find the infinite loop and value of the accumulator
	visited := make(map[int]bool)
	jumps := make([]int, 0)
	noops := make([]int, 0)
	index := 0
	acc := 0
	for true {
		if visited[index] {
			break
		}
		visited[index] = true
		if strings.Contains(instructions[index], "jmp") {
			jumps = append(jumps, index)
		} else if strings.Contains(instructions[index], "noop"){
			noops = append(noops, index)
		}
		index, acc = process(instructions[index], index, acc)
	}
	fmt.Println("Problem 1 solution:")
	fmt.Println(acc)
	fmt.Println("Problem 2 solution:")
	loopy := true
	for _, noop := range noops {
		instructions[noop] = strings.Replace(instructions[noop], "nop", "jmp", 1)
		acc, loopy = hasLoop(instructions)
		if !loopy {
			fmt.Println(acc)
		}
		//fix what we messed up
		instructions[noop] = strings.Replace(instructions[noop], "jmp", "nop", 1)
	}
	for _, jump := range jumps {
		instructions[jump] = strings.Replace(instructions[jump], "jmp", "nop", 1)
		acc, loopy = hasLoop(instructions)
		if !loopy {
			fmt.Println(acc)
		}
		//fix what we messed up
		instructions[jump] = strings.Replace(instructions[jump], "nop", "jmp", 1)
	}

}

func hasLoop(instructions []string) (int, bool) {
	visited := make(map[int]bool)
	acc := 0
	index := 0
	for true {
                if visited[index] {
			return acc, true
                }
                visited[index] = true
                index, acc = process(instructions[index], index, acc)
		if index >= len(instructions) {
			break
		}
        }
	return acc, false
}

func process(instruction string, index int, acc int) (int, int) {
	tokens := strings.Split(instruction, " ")
	switch tokens[0] {
	case "jmp":
		jump, _ := strconv.Atoi(tokens[1])
		index = index + jump
	case "acc":
		increment, _ := strconv.Atoi(tokens[1])
		acc = acc + increment
		index = index + 1
	case "nop":
		//do nothing
		index = index + 1
	}
	return index, acc
}
