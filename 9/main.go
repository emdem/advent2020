package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	fmt.Println("Part 1 answer:")
	index := 0
	solution := 0
	for true {
		if index + 26 > len(numbers) {
			break
		}
		if !properSum(numbers[index:index+26]) {
			solution = numbers[index+25]
			fmt.Println(solution)
			break
		}
		index = index + 1
	}
	fmt.Println("Part 2 answer:")
	slice_size := 2
	index = 0
	for true {
		if slice_size > 1024 { break }
		checkSum := sum(numbers[index:index+slice_size])
		if slice_size+index >= len(numbers) || numbers[index+slice_size] > solution {
			index = 0
			slice_size = slice_size + 1
		}
		if checkSum == solution {
			solutionSlice := numbers[index:index+slice_size]
			sort.Ints(solutionSlice)
			fmt.Println(solutionSlice[0]+solutionSlice[slice_size-1])
			break
		}
		index = index + 1
	}
}

func sum(slice []int) int {
	result := 0
	for _, v := range slice {
		result = result + v
	}
	return result
}

func properSum(slice []int) bool {
	if len(slice) != 26 {
		fmt.Println("This slice is too small foo")
		return false
	}
	for i, v := range slice[0:25-1] {
		for _, w := range slice[i+1:25] {
			if v != w && v+w==slice[25] {
				return true
			}
		}
	}
	return false
}
