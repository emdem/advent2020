package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/deckarep/golang-set"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sets []mapset.Set
	unionCount := 0
	intersectionCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if len(sets) != 0 {
				unionSet := mapset.NewSet()
				intersectSet := mapset.NewSet()
				for i, s := range sets {
					if i == 0 {
						intersectSet = intersectSet.Union(s)
					}
					unionSet = unionSet.Union(s)
					intersectSet = intersectSet.Intersect(s)
				}
				unionCount = unionCount + unionSet.Cardinality()
				intersectionCount = intersectionCount + intersectSet.Cardinality()
			}
			sets = make([]mapset.Set, 0)
		} else {
			set := mapset.NewSet()
			for _, c := range line {
				set.Add(c)
			}
			sets = append(sets, set)
		}
	}
	//Problem 1 - count unique yeses total by group
	// union of all answers per group
	fmt.Println("Problem 1 answer:")
	fmt.Println(unionCount)
	//Problem 2 - intersection of all answers per group
	fmt.Println("Problem 2 answer:")
	fmt.Println(intersectionCount)
}
