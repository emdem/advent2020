package main

import (
	"bufio"
	"fmt"
	"log"
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
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, x)
	}
	for i, s1 := range result {
		for j, s2 := range result {
			if i == j {
				// do nothing
			}
			if s1 + s2 == 2020 {
				fmt.Println("Answer 1:")
				fmt.Println(s1*s2)
			}
		}
	}
	for i, s1 := range result {
		for j, s2 := range result {
			for k, s3 := range result {
				if i == j || j == k || i == k {
					// do nothing
				}
				if s1 + s2 + s3 == 2020 {
					fmt.Println("Answer 2:")
					fmt.Println(s1*s2*s3)
				}
			}
		}
	}
}
