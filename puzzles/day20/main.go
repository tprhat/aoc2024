package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

type QItem struct {
	p    Point
	dist int
}

func bfs(matrix [][]string, pnt Point) map[Point]int {
	Q := []QItem{{pnt, 0}}
	dists := make(map[Point]int)

	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]
		if _, exists := dists[u.p]; exists {
			continue
		}
		if matrix[u.p.x][u.p.y] != "#" {
			dists[u.p] = u.dist
		}
		for _, dir := range dirs {
			nx, ny := u.p.x+dir.dx, u.p.y+dir.dy
			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] != "#" {
				Q = append(Q, QItem{Point{nx, ny}, u.dist + 1})
			}
		}

	}

	return dists
}
func dijkstra(matrix [][]string, source, end Point) int {
	var Q []QueueItem
	visited := []Point{}
	Q = append(Q, QueueItem{source.x, source.y, 0})
	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]

		if u.x == end.x && u.y == end.y {
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

// the idea behind part 1 was to remove one wall from the whole grid and run path finding from S to E for each wall removed
// this worked but it was extremely slow (5m30s) so another plan had to be found.
// the plan for part 2 was to create 2 maps of all the distances from start and end to every point in the matrix.
// the algorithm goes over all the points it the starting and ending map together and checks if the distance
// between 2 points from start and end was <= 20, if that is fine it checks if it's at least a 100 better than the best solution
// in the end it turned out that it works for any distance so part1 uses the same algorithm.
func solve(name string) (int, int) {
	matrix, start, end := transformInput(name)
	best := dijkstra(matrix, start, end)
	part1, part2 := 0, 0
	fromstart := bfs(matrix, start)
	fromend := bfs(matrix, end)
	for pstart, distStart := range fromstart {
		for pEnd, distEnd := range fromend {
			pntDist := math.Abs(float64(pstart.x)-float64(pEnd.x)) + math.Abs(float64(pstart.y)-float64(pEnd.y))
			if pntDist <= 2 {
				if distStart+distEnd+int(pntDist) <= best-100 {
					part1 += 1
				}
			}
			if pntDist <= 20 {
				if distStart+distEnd+int(pntDist) <= best-100 {
					part2 += 1
				}
			}
		}
	}
	return part1, part2
}

func transformInput(name string) ([][]string, Point, Point) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	lines := [][]string{}
	scanner := bufio.NewScanner(file)
	i := 0
	var start, end Point
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for j, l := range line {
			if l == "S" {
				start = Point{i, j}
			}
			if l == "E" {
				end = Point{i, j}
			}
		}
		lines = append(lines, line)
		i++
	}
	return lines, start, end
}

func main() {
	start := time.Now()
	part1, part2 := solve("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	t1 := time.Now()
	fmt.Println("Total time: ", t1.Sub(start))
}
