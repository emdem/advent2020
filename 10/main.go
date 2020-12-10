package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"sort"
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
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	sort.Ints(numbers)
	gaps := make(map[int]int)
	fmt.Println("Answer to part 1:")
	//first one is always 0->#
	gaps[numbers[0]] = gaps[numbers[0]]+1
	//second problem wants permutations
	permutations := 1
	singles := make(map[int]bool)
	for i, _ := range numbers {
		difference := 3 //for the last one
		if i + 1 < len(numbers) {
			difference = numbers[i+1]-numbers[i]
		}
		gaps[difference] = gaps[difference] + 1
		if difference == 1 {
			singles[numbers[i]] = true
		} else if difference > 1 {
			if len(singles) > 1 {
				permutations = permutations * int(math.Pow(2, float64(len(singles)-1)) - math.Max(float64(len(singles)-3), 0.0))}
			singles = make(map[int]bool)
		}
	}
	fmt.Println(gaps[1]*gaps[3])
	fmt.Println("Answer to part 2:")
	fmt.Println(permutations*2) //*2 is so lazy, but meh... the math checks out
}
