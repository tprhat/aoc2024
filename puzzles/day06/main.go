package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1:", part1("input.txt"))
	t1 := time.Now()
	fmt.Println("Part 1 time: ", t1.Sub(start))
	fmt.Println("Part 2:", part2("input.txt"))
	t2 := time.Now()
	fmt.Println("Part 2 time: ", t2.Sub(t1))
}

func part1(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	pos := []int{}
	for i, line := range lines {
		for j, l := range line {
			if string(l) == "^" {
				pos = append(pos, i, j)
			}
		}
		matrix = append(matrix, strings.Split(line, ""))
	}
	dir := []int{-1, 0}

	visited := make(map[string]int)
	visited[toString(pos)] += 1
	for {
		// if outside of the board
		if pos[0]+dir[0] < 0 || pos[0]+dir[0] >= len(matrix) || pos[1]+dir[1] < 0 || pos[1]+dir[1] >= len(matrix) {
			break
		}
		if matrix[pos[0]+dir[0]][pos[1]+dir[1]] == "#" {
			dir = turnRight(dir)
			continue
		}
		pos[0] += dir[0]
		pos[1] += dir[1]
		visited[toString(pos)] += 1
	}
	return len(visited)
}

func toString(pos []int) string {
	return strconv.Itoa(pos[0]) + "," + strconv.Itoa(pos[1])
}

func turnRight(pos []int) []int {
	// up [-1, 0]
	// right [0, 1]
	// down [-1, 0]
	// left [0, -1]
	x := 0
	y := 0
	if pos[0] != 0 {
		y = pos[0] * -1
	}
	if pos[1] != 0 {
		x = pos[1]
	}
	return []int{x, y}
}

type Position struct {
	x int
	y int
}

func part2(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	posStarting := []int{}
	for i, line := range lines {
		for j, l := range line {
			if string(l) == "^" {
				posStarting = append(posStarting, i, j)
			}
		}
		matrix = append(matrix, strings.Split(line, ""))
	}
	// try to put a box on every position in the visited list and see if it contains a loop
	cnt := 0
	visitedPoints := getVisitedPoints(matrix, posStarting)
	for _, point := range visitedPoints {
		// write the starting position back to pos on every iteration
		// since it gets changed in the inner loop
		pos := Position{x: posStarting[0], y: posStarting[1]}
		if point.x == pos.x && point.y == pos.y {
			// skip starting position
			continue
		}
		matrix[point.x][point.y] = "#"
		dir := []int{-1, 0}
		dirNum := 1
		visited := make(map[Position]int)
		visited[pos] = dirNum
		for {
			// if outside of the board
			if pos.x+dir[0] < 0 || pos.x+dir[0] >= len(matrix) || pos.y+dir[1] < 0 || pos.y+dir[1] >= len(matrix) {
				break
			}
			if matrix[pos.x+dir[0]][pos.y+dir[1]] == "#" {
				dirNum = (dirNum)%4 + 1
				dir = turnRight(dir)
				continue
			}
			pos.x += dir[0]
			pos.y += dir[1]

			// found a loop
			if visited[pos] == dirNum {
				cnt++
				break
			}
			visited[pos] = dirNum
		}
		matrix[point.x][point.y] = "."
	}
	return cnt
}

func getVisitedPoints(matrix [][]string, poss []int) []Position {
	dir := []int{-1, 0}
	pos := []int{}
	pos = append(pos, poss...)
	var visitedPoints []Position
	// Position{} creates a new pointer
	visitedPoints = append(visitedPoints, Position{x: pos[0], y: pos[1]})
	for {
		// if outside of the board
		if pos[0]+dir[0] < 0 || pos[0]+dir[0] >= len(matrix) || pos[1]+dir[1] < 0 || pos[1]+dir[1] >= len(matrix) {
			break
		}
		if matrix[pos[0]+dir[0]][pos[1]+dir[1]] == "#" {
			dir = turnRight(dir)
			continue
		}
		pos[0] += dir[0]
		pos[1] += dir[1]
		exists := false
		for _, elem := range visitedPoints {
			if elem.x == pos[0] && elem.y == pos[1] {
				exists = true
				break
			}
		}
		if !exists {
			visitedPoints = append(visitedPoints, Position{x: pos[0], y: pos[1]})
		}
	}
	return visitedPoints
}

func transformInput(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
