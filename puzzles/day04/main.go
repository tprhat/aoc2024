package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))

}

func part1(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	count := 0
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			if matrix[i][j] == "X" {
				count += findWord(matrix, i, j)
			}
		}
	}

	return count
}

func findWord(matrix [][]string, i int, j int) int {
	count := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	for _, dir := range directions {
		if checkDirection(matrix, i, j, dir) {
			count++
		}
	}

	return count
}
func checkDirection(matrix [][]string, i int, j int, direction []int) bool {
	// since the "X" was found we only check for the remaining 3 characters
	word := []string{"M", "A", "S"}
	i_new := i
	j_new := j
	for _, w := range word {
		i_new, j_new = i_new+direction[0], j_new+direction[1]

		if i_new < 0 || j_new < 0 || i_new >= len(matrix) || j_new >= len(matrix[0]) {
			return false
		}
		if matrix[i_new][j_new] != w {
			return false
		}
	}
	return true
}

func part2(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	count := 0
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			if matrix[i][j] == "A" {
				if findWord2(matrix, i, j) {
					count++
				}
			}
		}
	}

	return count
}

func findWord2(matrix [][]string, i, j int) bool {
	boolDiag1 := false
	boolDiag2 := false

	if i-1 < 0 || i+1 >= len(matrix) || j-1 < 0 || j+1 >= len(matrix[0]) {
		return false
	}

	if ((matrix[i-1][j-1] == "M") && (matrix[i+1][j+1] == "S")) || ((matrix[i-1][j-1] == "S") && (matrix[i+1][j+1] == "M")) {
		boolDiag1 = true
	}
	if ((matrix[i-1][j+1] == "M") && (matrix[i+1][j-1] == "S")) || ((matrix[i-1][j+1] == "S") && (matrix[i+1][j-1] == "M")) {
		boolDiag2 = true
	}

	if boolDiag1 && boolDiag2 {
		return true
	}
	return false

}

func transformInput(name string) []string {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")

	return lines

}
