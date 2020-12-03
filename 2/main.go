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
	goodPassCount1 := 0
	goodPassCount2 := 0
	for scanner.Scan() {
		//SO GROSS
		tokens := strings.Split(scanner.Text(), ":")
		password := tokens[1]
		tokens = strings.Split(tokens[0], " ")
		letter := tokens[1]
		tokens = strings.Split(tokens[0], "-")
		lowCount, _ := strconv.Atoi(tokens[0])
		highCount, _ := strconv.Atoi(tokens[1])
		count := strings.Count(password, letter)
		//for the first problem
		if count >= lowCount && count <= highCount {
			goodPassCount1 = goodPassCount1 + 1
		}
		//for the second problem
		if (password[lowCount] == letter[0]) != (password[highCount] == letter[0]) {
			goodPassCount2 = goodPassCount2 + 1
		}
	}
	fmt.Println("part one answer:")
	fmt.Println(goodPassCount1)
	fmt.Println("part two answer:")
	fmt.Println(goodPassCount2)
}
