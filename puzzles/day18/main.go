package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}
type QueueItem struct {
	x, y, dist int
}

var dirs []Direction = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func dijkstra(matrix [][]string, source Point) int {
	var Q []QueueItem
	visited := []Point{}
	Q = append(Q, QueueItem{source.x, source.y, 0})
	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]

		if u.x == 70 && u.y == 70 {
			return u.dist
		}
		if slices.Contains(visited, Point{u.x, u.y}) {
			continue
		}
		visited = append(visited, Point{u.x, u.y})

		for _, dir := range dirs {
			nx, ny := u.x+dir.dx, u.y+dir.dy
			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] != "#" {
				Q = append(Q, QueueItem{nx, ny, u.dist + 1})
			}
		}
	}
	return -1
}

func part1(name string) int {
	walls, _ := transformInput(name, 1024)
	matrix := make([][]string, 71)
	for i := range 71 {
		matrix[i] = make([]string, 71)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, exists := walls[Point{i, j}]; exists {
				matrix[i][j] = "#"
			} else {
				matrix[i][j] = "."
			}
		}
	}
	return dijkstra(matrix, Point{0, 0})
}

func part2(name string) string {
	var strLast string
	// the 3000 came from trying things out
	for n := 3000; n <= 3450; n++ {
		walls, lastCoord := transformInput(name, n)
		matrix := make([][]string, 71)
		for i := range 71 {
			matrix[i] = make([]string, 71)
		}
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if _, exists := walls[Point{i, j}]; exists {
					matrix[i][j] = "#"
				} else {
					matrix[i][j] = "."
				}
			}
		}
		shortestPath := dijkstra(matrix, Point{0, 0})
		if shortestPath == -1 {
			strLast = lastCoord
			break
		}
	}
	return strLast
}

func transformInput(name string, nWalls int) (map[Point]bool, string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	walls := make(map[Point]bool)
	scanner := bufio.NewScanner(file)
	i := 0
	str := ""
	for scanner.Scan() {
		if i == nWalls {
			break
		}
		str = scanner.Text()
		s := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(s[0])
		x, _ := strconv.Atoi(s[1])
		walls[Point{x, y}] = true
		i++
	}
	return walls, str
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
