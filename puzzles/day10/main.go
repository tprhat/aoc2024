package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// the first part would be done in 5 minutes
// if I used global variables from the beginning
var count int = 0
var visitedTops [][]int

func part1(name string) int {
	matrix := transformInput(name)
	for i, line := range matrix {
		for j, point := range line {
			if point == 0 {
				countTops(matrix, i, j, 1, true)
				visitedTops = [][]int{}
			}
		}
	}
	return count
}

func part2(name string) int {
	count = 0
	visitedTops = [][]int{}
	matrix := transformInput(name)
	for i, line := range matrix {
		for j, point := range line {
			if point == 0 {
				countTops(matrix, i, j, 1, false)
				visitedTops = [][]int{}
			}
		}
	}
	return count
}

// the idea is basically to run a DFS across the matrix when it finds a 0
// the problem was that I wanted to avoid global variable and kept getting stuck
// with counting when the top was reached
// part 1 wanted distinct number of tops reached
// part 2 wanted non-distinct number of tops reached
// that's why there is the isPart1 bool :D
func countTops(matrix [][]int, i, j, next int, isPart1 bool) {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for _, d := range dirs {
		if i+d[0] < 0 || i+d[0] >= len(matrix) || j+d[1] < 0 || j+d[1] >= len(matrix) {
			continue
		}
		if matrix[i+d[0]][j+d[1]] == next {
			if next == 9 {
				if isPart1 {
					isVisited := false
					for _, visited := range visitedTops {
						if i+d[0] == visited[0] && j+d[1] == visited[1] {
							isVisited = true
						}
					}
					if isVisited {
						continue
					}
				}
				// found a top
				visitedTops = append(visitedTops, []int{i + d[0], j + d[1]})
				count++
			}
			countTops(matrix, i+d[0], j+d[1], next+1, isPart1)
		}
	}
}

func transformInput(name string) [][]int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	matrix := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}
		for _, elem := range scanner.Text() {

			line = append(line, int(elem-'0'))
		}
		matrix = append(matrix, line)
	}

	return matrix
}

func main() {
	start := time.Now()
	fmt.Println("Part 1:", part1("input.txt"))
	t1 := time.Now()
	fmt.Println("Part1 time: ", t1.Sub(start))
	fmt.Println("Part 2:", part2("input.txt"))
	t2 := time.Now()
	fmt.Println("Part2 time: ", t2.Sub(t1))
}
