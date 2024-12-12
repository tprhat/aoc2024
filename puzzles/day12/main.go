package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

var seen [][]int = [][]int{}

func part1(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			area, perimiter := dfs(matrix, matrix[i][j], i, j, 0, 0)
			total += area * perimiter
		}
	}
	return total
}

func dfs(matrix [][]string, region string, i, j, area, perimiter int) (int, int) {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	if isSeen(i, j) {
		return 0, 0
	}
	// perimeter = 4 - elements in all 4 directions the same type
	// check how many fences we can add for this position
	maxper := 4
	for _, dir := range dirs {
		if i+dir[0] < 0 || i+dir[0] >= len(matrix) || j+dir[1] < 0 || j+dir[1] >= len(matrix) {
			continue
		}
		if matrix[i+dir[0]][j+dir[1]] == region {
			maxper -= 1
		}
	}
	seen = append(seen, []int{i, j})
	atot, ptot := 1, maxper
	for _, dir := range dirs {
		if i+dir[0] < 0 || i+dir[0] >= len(matrix) || j+dir[1] < 0 || j+dir[1] >= len(matrix) {
			continue
		}
		if matrix[i+dir[0]][j+dir[1]] == region {
			a, p := dfs(matrix, region, i+dir[0], j+dir[1], area+1, perimiter+maxper)
			atot += a
			ptot += p
		}
	}
	return atot, ptot
}

func isSeen(i, j int) bool {
	for _, elem := range seen {
		if elem[0] == i && elem[1] == j {
			return true
		}
	}
	return false
}

type Point struct {
	x, y int
}

type Region struct {
	plantType string
	cells     []Point
	area      int
	sides     int
}

func part2(name string) int {
	var matrix [][]string
	lines := transformInput(name)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	seen = [][]int{}
	total := findRegions(matrix)
	return total
}

func findRegions(matrix [][]string) int {
	total := 0
	for i := range matrix {
		for j := range matrix {
			if isSeen(i, j) {
				continue
			}
			region := &Region{
				plantType: matrix[i][j],
			}
			fillRegion(matrix, i, j, region)

			region.area = len(region.cells)
			region.sides = countSides(matrix, region)

			total += region.area * region.sides
		}
	}
	return total
}

func fillRegion(matrix [][]string, i, j int, region *Region) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix) ||
		isSeen(i, j) || matrix[i][j] != region.plantType {
		return
	}

	seen = append(seen, []int{i, j})
	region.cells = append(region.cells, Point{x: i, y: j})

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for _, dir := range dirs {
		fillRegion(matrix, i+dir[0], j+dir[1], region)
	}
}

func countSides(matrix [][]string, region *Region) int {
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	perimeterLines := [][]Point{}
	for _, cell := range region.cells {
		for _, dir := range dirs {
			newX, newY := cell.x+dir[0], cell.y+dir[1]
			// we can represent a perimeter line by having 2 points and the line is "between" them
			// the second point must not be a part of that region
			if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix) ||
				!slices.Contains(region.cells, Point{x: newX, y: newY}) {
				perimeterLines = append(perimeterLines, []Point{cell, {newX, newY}})
			}
		}
	}
	dirs = [][]int{
		{0, 1},
		{1, 0},
	}

	// HARD PART OF THE ALGORITHM
	sides := [][]Point{}
	for _, p := range perimeterLines {
		keep := true
		// only move in two directions up and right
		// since we don't want two element touching diagonally
		// to count as one side
		// BB..
		// ---- --> this should be 2 sides
		// ..BB
		for _, dir := range dirs {
			p1n := Point{p[0].x + dir[0], p[0].y + dir[1]}
			p2n := Point{p[1].x + dir[0], p[1].y + dir[1]}
			// if the moved line is in the perimeter lines
			// we want to ignore it since we only want one element per side
			if containsPoint(perimeterLines, p1n, p2n) {
				keep = false
				break
			}
		}
		if keep {
			sides = append(sides, []Point{p[0], p[1]})
		}
	}
	return len(sides)
}

func containsPoint(per [][]Point, p1n, p2n Point) bool {
	for _, p := range per {
		if p[0] == p1n && p[1] == p2n {
			return true
		}
	}
	return false
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

func main() {
	start := time.Now()
	fmt.Println("Part 1:", part1("input.txt"))
	t1 := time.Now()
	fmt.Println("Part1 time: ", t1.Sub(start))
	fmt.Println("Part 2:", part2("input.txt"))
	t2 := time.Now()
	fmt.Println("Part2 time: ", t2.Sub(t1))
}
