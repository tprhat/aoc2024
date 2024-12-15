package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

var dirs map[string]Point = map[string]Point{"<": {0, -1}, ">": {0, 1}, "^": {-1, 0}, "v": {1, 0}}

func part1(name string) int {
	lines, moves := transformInput(name)
	var matrix [][]string
	var currPoint Point
	for i, line := range lines {
		l := strings.Split(line, "")
		for j, elem := range l {
			if elem == "@" {
				currPoint = Point{i, j}
			}
		}
		matrix = append(matrix, l)
	}
	// fmt.Println(currPoint, matrix, moves)
	for _, move := range moves {
		matrix, currPoint = moveInDir(matrix, currPoint, move)
		// fmt.Println(move)
		// fmt.Println(currPoint)
		// printMap(matrix)
	}
	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "O" {
				total += 100*i + j
			}
		}
	}
	return total
}
func printMap(matrix [][]string) {
	fmt.Println("_____MAP_______")
	for _, m := range matrix {
		fmt.Println(m)
	}
	fmt.Println("____END________")
}
func moveInDir(matrix [][]string, pos Point, dir string) ([][]string, Point) {
	currPos := pos
	currPosRobot := pos
	nStones := 0
	currDir := dirs[dir]
	for {
		currPos.x += currDir.x
		currPos.y += currDir.y
		// fmt.Println("DEBUG1", currPos)
		// fmt.Println(matrix[currPos.x][currPos.y])
		if matrix[currPos.x][currPos.y] == "#" {
			nStones = 0
			break
		}
		if matrix[currPos.x][currPos.y] == "." {
			nStones++
			break
		}
		if matrix[currPos.x][currPos.y] == "O" {
			nStones++
		}
	}
	// reset currPos to starting
	if nStones > 0 {
		currPosRobot = pos
		matrix[currPosRobot.x][currPosRobot.y] = "."
		currPosRobot.x += currDir.x
		currPosRobot.y += currDir.y
		matrix[currPosRobot.x][currPosRobot.y] = "@"
		// fmt.Println("ROBOT MOVED TO:", currPosRobot)
	}
	if nStones > 1 {
		currPos = currPosRobot
		// fmt.Println("MOVING ROCKS", currPos) // OK
		// fmt.Println(currPos)
		for i := 0; i < nStones-1; i++ {
			currPos.x += currDir.x
			currPos.y += currDir.y
			matrix[currPos.x][currPos.y] = "O"
		}
	}
	return matrix, currPosRobot
}

func part2(name string) int {
	lines, moves := transformInput(name)
	var matrix [][]string
	var currPoint Point
	for i, line := range lines {
		l := strings.Split(line, "")
		l2 := []string{}
		for j, elem := range l {
			if elem == "@" {
				currPoint = Point{i, j}
				l2 = append(l2, "@", ".")
			} else if elem == "#" {
				l2 = append(l2, "#", "#")
			} else if elem == "." {
				l2 = append(l2, ".", ".")
			} else if elem == "O" {
				l2 = append(l2, "[", "]")
			}
		}
		matrix = append(matrix, l2)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "@" {
				currPoint = Point{i, j}
			}
		}
	}
	for _, move := range moves {
		matrix, currPoint = moveInDir2(matrix, currPoint, move)
	}

	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "[" {
				total += 100*i + j
			}
		}
	}
	return total
}

type Point2 struct {
	x, y int
	sign string
}

func moveInDir2(matrix [][]string, pos Point, dir string) ([][]string, Point) {
	currPos := pos
	currPosRobot := pos
	nStones := 0
	currDir := dirs[dir]
	if dir == "<" || dir == ">" {
		for {
			currPos.x += currDir.x
			currPos.y += currDir.y
			if matrix[currPos.x][currPos.y] == "[" || matrix[currPos.x][currPos.y] == "]" {
				nStones++
			}
			if matrix[currPos.x][currPos.y] == "." {
				nStones++
				break
			}
			if matrix[currPos.x][currPos.y] == "#" {
				nStones = 0
				break
			}
		}
		var nextElem string
		if nStones > 0 {
			currPosRobot = pos
			matrix[currPosRobot.x][currPosRobot.y] = "."
			currPosRobot.x += currDir.x
			currPosRobot.y += currDir.y
			nextElem = matrix[currPosRobot.x][currPosRobot.y]
			matrix[currPosRobot.x][currPosRobot.y] = "@"
			// fmt.Println("ROBOT MOVED TO:", currPosRobot)
		}
		if nStones > 1 {
			currPos = currPosRobot
			// fmt.Println("MOVING ROCKS", currPos) // OK
			// fmt.Println(currPos)
			for i := 0; i < nStones-1; i++ {
				currPos.x += currDir.x
				currPos.y += currDir.y

				matrix[currPos.x][currPos.y] = nextElem
				if nextElem == "[" {
					nextElem = "]"
				} else {
					nextElem = "["
				}
			}
		}
		return matrix, currPosRobot

	}
	// up and down
	stonesToMove := [][]Point2{}
	levels := 0
	for {
		currLevelStones := []Point2{}
		currPos.x += currDir.x
		currPos.y += currDir.y
		if len(stonesToMove) == 0 {
			if matrix[currPos.x][currPos.y] == "." {
				levels++
				break
			}
			if matrix[currPos.x][currPos.y] == "#" {
				levels = 0
				break
			}
			if matrix[currPos.x][currPos.y] == "[" {
				levels++
				currLevelStones = append(currLevelStones, Point2{currPos.x, currPos.y, "["}, Point2{currPos.x, currPos.y + 1, "]"})
			}
			if matrix[currPos.x][currPos.y] == "]" {
				levels++
				currLevelStones = append(currLevelStones, Point2{currPos.x, currPos.y - 1, "["}, Point2{currPos.x, currPos.y, "]"})
			}
			stonesToMove = append(stonesToMove, currLevelStones)
			continue
		}
		// [4,6 4,7]
		for _, previousLevelStones := range stonesToMove[levels-1] {
			currPos = Point{previousLevelStones.x + currDir.x, previousLevelStones.y + currDir.y}
			if matrix[currPos.x][currPos.y] == "#" {
				levels = 0
				break
			}
			if matrix[currPos.x][currPos.y] == "." {
				continue
			}
			if matrix[currPos.x][currPos.y] == "[" {
				if !checkIfContains(currLevelStones, currPos) {
					currLevelStones = append(currLevelStones, Point2{currPos.x, currPos.y, "["}, Point2{currPos.x, currPos.y + 1, "]"})
				}
			}
			if matrix[currPos.x][currPos.y] == "]" {
				if !checkIfContains(currLevelStones, currPos) {
					currLevelStones = append(currLevelStones, Point2{currPos.x, currPos.y - 1, "["}, Point2{currPos.x, currPos.y, "]"})
				}
			}

		}
		if levels == 0 {
			return matrix, pos
		}
		levels++
		if len(currLevelStones) == 0 {
			break
		}
		stonesToMove = append(stonesToMove, currLevelStones)

	}
	if levels == 0 {
		return matrix, pos

	}
	for j := len(stonesToMove) - 1; j >= 0; j-- {
		for _, stone := range stonesToMove[j] {
			matrix[stone.x][stone.y] = "."
			matrix[stone.x+currDir.x][stone.y+currDir.y] = stone.sign
		}
	}
	currPosRobot = pos
	matrix[currPosRobot.x][currPosRobot.y] = "."
	currPosRobot.x += currDir.x
	currPosRobot.y += currDir.y
	matrix[currPosRobot.x][currPosRobot.y] = "@"

	return matrix, currPosRobot

}

func checkIfContains(lst []Point2, pos Point) bool {
	for _, pnt := range lst {
		if pnt.x == pos.x && pnt.y == pos.y {
			return true
		}
	}
	return false
}

func transformInput(name string) ([]string, []string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	lines := []string{}
	moves := []string{}
	isMap := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			isMap = false
			continue
		}
		if isMap {
			lines = append(lines, scanner.Text())
		} else {
			// expand a string to create a list of string elements with len of 1
			// instead of a slice of string elements with len > 1
			moves = append(moves, strings.Split(scanner.Text(), "")...)
		}
	}
	return lines, moves
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
