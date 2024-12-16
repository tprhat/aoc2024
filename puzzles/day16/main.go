package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}
type Maze struct {
	matrix [][]string
	start  Point
	end    Point
}
type QueueItem struct {
	pos   Point
	dir   int
	score int

	path []Point // used for part 2
}

var dirs = []Direction{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

var maze Maze
var lowestScore int

func part1(name string) int {
	lines := transformInput(name)
	var matrix [][]string
	var end Point
	var start Point
	for i, line := range lines {
		l := strings.Split(line, "")
		for j, elem := range l {
			if elem == "S" {
				start = Point{i, j}
			}
			if elem == "E" {
				end = Point{i, j}
			}
		}
		matrix = append(matrix, l)
	}
	maze = Maze{matrix: matrix, start: start, end: end}
	lowestScore = findLowestScore(maze)
	return lowestScore
}

// creates a key for visited map
func (p Point) key(dir int) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, dir)
}

func (p Point) add(d Direction) Point {
	return Point{p.x + d.dx, p.y + d.dy}
}

// check if new point is not a wall or out of bounds
func (m Maze) isValid(p Point) bool {
	return p.x >= 0 && p.x < len(m.matrix) &&
		p.y >= 0 && p.y < len(m.matrix[0]) &&
		m.matrix[p.x][p.y] != "#"
}

func findLowestScore(maze Maze) int {
	// heapq
	queue := []QueueItem{{maze.start, 0, 0, nil}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		// sort the queue to take the lowest scores first
		sort.Slice(queue, func(i, j int) bool { return queue[i].score < queue[j].score })

		current := queue[0]
		queue = queue[1:]

		if current.pos.x == maze.end.x && current.pos.y == maze.end.y {
			return current.score
		}
		key := current.pos.key(current.dir)
		if visited[key] {
			continue
		}
		visited[key] = true

		nextPos := current.pos.add(dirs[current.dir])
		// add forward move
		if maze.isValid(nextPos) {
			queue = append(queue, QueueItem{
				pos:   nextPos,
				dir:   current.dir,
				score: current.score + 1,
				path:  nil,
			})
		}
		// add right and left turn
		queue = append(queue,
			QueueItem{pos: current.pos, dir: (current.dir + 1) % 4, score: current.score + 1000, path: nil},
			QueueItem{pos: current.pos, dir: (current.dir + 3) % 4, score: current.score + 1000, path: nil},
		)
	}
	// if we ever come to here something is wrong
	return -1
}

// part 2 uses the created maze and lowestScore	from part 1
func part2() int {
	paths := findOptimalPaths(maze, lowestScore)
	return countUniquePointsInPaths(paths)
}

func findOptimalPaths(maze Maze, targetScore int) [][]Point {
	// almost the same thing as part 1, just save the visited point in the path
	queue := []QueueItem{{maze.start, 0, 0, []Point{maze.start}}}
	visited := make(map[string]int)
	var paths [][]Point

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > targetScore {
			continue
		}
		key := current.pos.key(current.dir)
		// check the visited score and if it's less than current.score we can move to the next path
		// since this is not optimal
		if score, exists := visited[key]; exists && score < current.score {
			continue
		}
		// write the score to visited
		visited[key] = current.score

		if current.pos.x == maze.end.x && current.pos.y == maze.end.y && current.score == targetScore {
			paths = append(paths, current.path)
			continue
		}

		nextPos := current.pos.add(dirs[current.dir])
		if maze.isValid(nextPos) {
			newPath := make([]Point, len(current.path))
			// copy the current path to a new object
			// otherwise it's going to change that memory location
			// when called elsewhere
			copy(newPath, current.path)
			queue = append(queue, QueueItem{
				pos:   nextPos,
				dir:   current.dir,
				score: current.score + 1,
				path:  append(newPath, nextPos),
			})
		}
		queue = append(queue,
			QueueItem{pos: current.pos, dir: (current.dir + 1) % 4, score: current.score + 1000, path: current.path},
			QueueItem{pos: current.pos, dir: (current.dir + 3) % 4, score: current.score + 1000, path: current.path},
		)
	}
	return paths
}

func countUniquePointsInPaths(paths [][]Point) int {
	uniqueP := make(map[string]bool)
	for _, path := range paths {
		for _, p := range path {
			uniqueP[p.key(0)] = true
		}
	}
	return len(uniqueP)
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
	fmt.Println("Part 2:", part2())
	t2 := time.Now()
	fmt.Println("Part2 time: ", t2.Sub(t1))
}
