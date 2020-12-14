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
	scanner := bufio.NewScanner(file)
	lines := make([]int, 0)
	rawLine := ""
	earliest := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			rawLine = line
			tokens := strings.Split(line, ",")
			for _, bus := range tokens {
				if bus != "x" {
					busLine, _ := strconv.Atoi(bus)
					lines = append(lines, busLine)
				}
			}
		} else {
			earliest, _ = strconv.Atoi(line)
		}
	}
	fmt.Println("Answer to part 1:")
	currentTime := earliest
	endLoop := false
	for true {
		for _, busLine := range lines {
			remainder := currentTime%busLine
			if remainder == 0 {
				fmt.Println((currentTime-earliest) * busLine)
				endLoop = true
				break;
			}
		}
		if endLoop {
			break;
		}
		currentTime = currentTime + 1
	}
	tokens := strings.Split(rawLine, ",")

	fmt.Println("Answer to part 2:")
	//ok, so brute force didn't work, lets make a map
	//of divisors to remainders
	divMap := make([][2]int, 0)
	for i, bus := range tokens {
		if bus != "x" {
			busLine, _ := strconv.Atoi(bus)
			remainder := (busLine-i)%busLine
			if remainder < 0 {
				remainder = remainder + busLine
			}
			divMap = append(divMap, [2]int{busLine, remainder})
		}
	}
	answer := 0
	inc := 1
	for _, pair := range divMap {
		for answer % pair[0] != pair[1] {
			answer = answer + inc
		}
		inc = inc * pair[0]
	}
	fmt.Println(answer)
}
