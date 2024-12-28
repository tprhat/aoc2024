package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"time"
)

var numericKeypad = map[rune]Point{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}

var directionalKeypad = map[rune]Point{
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

type Point struct {
	x, y int
}
type Direction struct {
	dx, dy int
}

var dirs = map[rune]Direction{
	'>': {0, 1},
	'<': {0, -1},
	'v': {1, 0},
	'^': {-1, 0},
}
var revDirectionKeypad = getReverseMap(directionalKeypad)
var revNumericKeypad = getReverseMap(numericKeypad)

var pairsMinDistanceCache map[string]int
var pathsCache map[string][]string

func solve(input []string, depth int) (res int) {
	for _, str := range input {
		// shortest sequence
		temp := getCost("A"+str, depth)
		// number from the task
		coeff, _ := strconv.Atoi(str[:len(str)-1])
		res += temp * coeff
	}
	return
}

func getReverseMap(m map[rune]Point) (w map[Point]rune) {
	w = make(map[Point]rune)
	for r, i := range m {
		w[i] = r
	}
	return
}

func getCost(str string, depth int) (res int) {
	for i := 0; i < len(str)-1; i++ {
		currPairCost := getPairCost(rune(str[i]), rune(str[i+1]), numericKeypad, revNumericKeypad, depth)
		res += currPairCost
	}
	return
}

func getPairCost(a, b rune, charToPoint map[rune]Point, pointToChar map[Point]rune, depth int) int {
	keyPadCode := 'd'
	// check if using numeric keypad
	if _, ok := charToPoint['0']; ok {
		keyPadCode = 'n'
	}
	key := fmt.Sprintf("%c%c%c%d", a, b, keyPadCode, depth)
	// if key exists return the distance from cache
	if dist, ok := pairsMinDistanceCache[key]; ok {
		return dist
	}
	// when we reach the final depth return the minimal distance from available paths
	if depth == 0 {
		minLen := math.MaxInt
		for _, path := range getAllPaths(a, b, directionalKeypad, revDirectionKeypad) {
			minLen = min(minLen, len(path))
		}
		return minLen
	}
	// find all paths
	allPaths := getAllPaths(a, b, charToPoint, pointToChar)
	minCost := math.MaxInt
	for _, path := range allPaths {
		path = "A" + path
		var currCost int
		for i := 0; i < len(path)-1; i++ {
			// recursively find path for each path from a to b
			currCost += getPairCost(rune(path[i]), rune(path[i+1]), directionalKeypad, revDirectionKeypad, depth-1)
		}
		minCost = min(minCost, currCost)
	}
	pairsMinDistanceCache[key] = minCost
	return minCost
}

func getAllPaths(a, b rune, charToPoint map[rune]Point, pointToChar map[Point]rune) (allPaths []string) {
	key := fmt.Sprintf("%c %c", a, b)
	// if path in cache return
	if paths, ok := pathsCache[key]; ok {
		return paths
	}
	DFS(charToPoint[a], charToPoint[b], []rune{}, charToPoint, pointToChar, make(map[Point]bool), &allPaths)
	pathsCache[key] = allPaths
	return
}

// dfs to get all paths from a to b
func DFS(curr, end Point, path []rune, charToPoint map[rune]Point, pointToChar map[Point]rune, visited map[Point]bool, allPaths *[]string) {
	if curr == end {
		*allPaths = append(*allPaths, string(path)+"A")
		return
	}
	visited[curr] = true
	for char, dir := range dirs {
		nIdx := Point{curr.x + dir.dx, curr.y + dir.dy}
		if _, ok := pointToChar[nIdx]; ok && !visited[nIdx] {
			newPath := slices.Clone(path)
			DFS(nIdx, end, append(newPath, char), charToPoint, pointToChar, visited, allPaths)
		}
	}
	visited[curr] = false
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
	pairsMinDistanceCache = make(map[string]int)
	pathsCache = make(map[string][]string)

	input := transformInput("input.txt")

	start := time.Now()
	fmt.Println("Part 1:", solve(input, 2))
	t1 := time.Now()
	fmt.Println("Part1 time: ", t1.Sub(start))
	fmt.Println("Part 2:", solve(input, 25))
	t2 := time.Now()
	fmt.Println("Part2 time: ", t2.Sub(t1))
}
