package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	bagMap := make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " bags contain ")
		color := tokens[0]
		contains := make(map[string]int)
		if tokens[1] != "no other bags." {
			tokenss := strings.Split(tokens[1], ", ")
			for _, tokez := range tokenss {
				tokenz := strings.Split(tokez, " ")
				contains[tokenz[1]+" "+tokenz[2]], _ = strconv.Atoi(tokenz[0])
			}
		}
		bagMap[color] = contains
	}

	fmt.Println("Answer 1:")
	fmt.Println(len(uniqueContains(bagMap, "shiny gold")))
	fmt.Println("Answer 2:")
	fmt.Println(totalContained(bagMap, "shiny gold"))
}

func uniqueContains(bagMap map[string]map[string]int, color string) map[string]bool {
	contains := make(map[string]bool)
	for k, v := range bagMap {
		for c := range v {
			if c == color {
				contains[k] = true
				for rec := range uniqueContains(bagMap, k) {
					contains[rec] = true
				}
			}
		}
	}
	return contains
}

func totalContained(bagMap map[string]map[string]int, color string) int {
	contained := bagMap[color]
	total := 0
	for color, count := range contained {
		total = total + count * (1 + totalContained(bagMap, color))
	}
	return total
}
