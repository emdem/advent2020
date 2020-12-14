package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func indexOf(a rune, dirs [4]rune) (int) {
	for k, v := range dirs {
		if a == v {
			return k
		}
	}
	return -1
}

func move(facing rune, magnitude int, x int, y int) (int, int) {
	switch facing {
	case 'N':
		y = y + magnitude
	case 'S':
		y = y - magnitude
	case 'E':
		x = x + magnitude
	case 'W':
		x = x - magnitude
	}
	return x, y
}

func left(turns, x, y int) (int, int) {
	//counter-clockwise
	//90 degrees (x,y) = (-y,x)
	if turns > 1 {
		return left(turns-1, -y, x)
	}
	return -y, x
}

func right(turns, x, y int) (int, int) {
	//clock-wise
	//90 degrees (x,y) = (y, -x)
	if turns > 1 {
		return right(turns-1, y, -x)
	}
	return y, -x
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("Answer to part 1:")
	x:=0 //(-)west<--->east(+)
	y:=0 //(-)south<--->north(+)
	dirs := [4]rune{'N','E','S','W'}
	facing := 'E'
	directions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		directions = append(directions, line)
		dir := line[0]
		dist, _ := strconv.Atoi(line[1:])
		if dir == 'L' || dir == 'R' {
			curInd := indexOf(rune(facing), dirs)
			dist = dist/90
			if dir == 'L' {
				curInd = curInd - dist
				if curInd < 0 {
					curInd = 4 + curInd
				}
			} else {
				curInd = (curInd + dist) % 4
			}
			facing = dirs[curInd]
		} else if dir == 'F' {
			x, y = move(facing, dist, x, y)
		} else {
			x, y = move(rune(dir), dist, x, y)
		}
	}
	fmt.Println(math.Abs(float64(x))+math.Abs(float64(y)))
	fmt.Println("Answer to part 2:")
	//waypoint vector
	wayX := 10 //(-)west<--->east(+)
	wayY := 1 //(-)south<--->north(+)
	x = 0 //ship vector
	y = 0
	for _, direction := range directions {
		dir := direction[0]
		dist, _ := strconv.Atoi(direction[1:])
		if dir == 'F' {
			//dist acts as a scalar
			x = x + (wayX * dist)
			y = y + (wayY * dist)
		} else if dir == 'L' {
			turns := dist/90
			wayX, wayY = left(turns, wayX, wayY)
		} else if dir == 'R' {
			turns := dist/90
			wayX, wayY = right(turns, wayX, wayY)
		} else {
			switch dir {
			case 'N':
				wayY = wayY + dist
			case 'S':
				wayY = wayY - dist
			case 'E':
				wayX = wayX + dist
			case 'W':
				wayX = wayX - dist
			}
			//ship don't move
		}
	}
	fmt.Println(math.Abs(float64(x))+math.Abs(float64(y)))
}
