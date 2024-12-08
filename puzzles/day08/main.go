package main

import (
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
	fmt.Println("Part1 time: ", t1.Sub(start))
	fmt.Println("Part 2:", part2("input.txt"))
	t2 := time.Now()
	fmt.Println("Part2 time: ", t2.Sub(t1))
}

func part1(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	antinodes := make(map[string]int)
	nodes := make(map[string][][]int)
	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] != "." {
				// key = station name, values = coordinates with that station name
				nodes[matrix[i][j]] = append(nodes[matrix[i][j]], []int{i, j})
				checkAntinodes(nodes, antinodes, matrix)
			}
		}
	}
	return len(antinodes)
}
func checkAntinodes(nodes map[string][][]int, antinodes map[string]int, matrix [][]string) {
	for _, v := range nodes {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				xdiff := v[j][0] - v[i][0]
				ydiff := v[j][1] - v[i][1]
				x1 := v[i][0] - xdiff
				y1 := v[i][1] - ydiff
				if x1 >= 0 && x1 < len(matrix) && y1 >= 0 && y1 < len(matrix) {
					antinodes[strconv.Itoa(x1)+","+strconv.Itoa(y1)] += 1
				}
				x2 := v[j][0] + xdiff
				y2 := v[j][1] + ydiff
				if x2 >= 0 && x2 < len(matrix) && y2 >= 0 && y2 < len(matrix) {
					antinodes[strconv.Itoa(x2)+","+strconv.Itoa(y2)] += 1
				}
			}
		}
	}
}

func part2(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	antinodes := make(map[string]int)
	nodes := make(map[string][][]int)
	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] != "." {
				nodes[matrix[i][j]] = append(nodes[matrix[i][j]], []int{i, j})
				antinodes[strconv.Itoa(i)+","+strconv.Itoa(j)] += 1
				checkAntinodes2(nodes, antinodes, matrix)
			}
		}
	}
	return len(antinodes)
}

func checkAntinodes2(nodes map[string][][]int, antinodes map[string]int, matrix [][]string) {
	for _, v := range nodes {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				xdiff := v[j][0] - v[i][0]
				ydiff := v[j][1] - v[i][1]

				x1 := v[i][0] - xdiff
				y1 := v[i][1] - ydiff
				// go in the same direction until the end of the matrix
				for x1 >= 0 && x1 < len(matrix) && y1 >= 0 && y1 < len(matrix) {
					if x1 >= 0 && x1 < len(matrix) && y1 >= 0 && y1 < len(matrix) {
						antinodes[strconv.Itoa(x1)+","+strconv.Itoa(y1)] += 1
					}
					x1 -= xdiff
					y1 -= ydiff
				}
				x2 := v[j][0] + xdiff
				y2 := v[j][1] + ydiff
				// go in the same direction until the end of the matrix
				for x2 >= 0 && x2 < len(matrix) && y2 >= 0 && y2 < len(matrix) {
					if x2 >= 0 && x2 < len(matrix) && y2 >= 0 && y2 < len(matrix) {
						antinodes[strconv.Itoa(x2)+","+strconv.Itoa(y2)] += 1
					}
					x2 += xdiff
					y2 += ydiff
				}
			}
		}
	}
}

func transformInput(name string) []string {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")
	return lines
}
