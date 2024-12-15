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
	for _, move := range moves {
		matrix, currPoint = moveInDir(matrix, currPoint, move)
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

func moveInDir(matrix [][]string, pos Point, dir string) ([][]string, Point) {
	currPos := pos
	currPosRobot := pos
	nStones := 0
	currDir := dirs[dir]
	for {
		currPos.x += currDir.x
		currPos.y += currDir.y
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
	}
	if nStones > 1 {
		currPos = currPosRobot
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

type Box struct {
	x, y int
	sign string
}

func moveInDir2(matrix [][]string, pos Point, dir string) ([][]string, Point) {
	currPos := pos
	currPosRobot := pos
	nStones := 0
	currDir := dirs[dir]
	// since the boxes now have a height of 1 and width of 2 split the code into 2 cases
	// left and right behave the same as in part 1 and up and down have a new logic
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
		}
		if nStones > 1 {
			currPos = currPosRobot
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
	// the idea was to add all the boxes and their coordinates to a slice of slices
	// that way the boxes at the same level are in the same slice
	// we can check the next levels movement based on all the boxes at the current level
	stonesToMove := [][]Box{}
	isEnd := false
	for {
		currLevelStones := []Box{}
		currPos.x += currDir.x
		currPos.y += currDir.y
		// first level
		if len(stonesToMove) == 0 {
			if matrix[currPos.x][currPos.y] == "." {
				break
			}
			if matrix[currPos.x][currPos.y] == "#" {
				isEnd = true
				break
			}
			if matrix[currPos.x][currPos.y] == "[" {
				currLevelStones = append(currLevelStones, Box{currPos.x, currPos.y, "["}, Box{currPos.x, currPos.y + 1, "]"})
			}
			if matrix[currPos.x][currPos.y] == "]" {
				currLevelStones = append(currLevelStones, Box{currPos.x, currPos.y - 1, "["}, Box{currPos.x, currPos.y, "]"})
			}
			stonesToMove = append(stonesToMove, currLevelStones)
			continue
		}
		// all levels after the first level
		for _, previousLevelStones := range stonesToMove[len(stonesToMove)-1] {
			currPos = Point{previousLevelStones.x + currDir.x, previousLevelStones.y + currDir.y}
			if matrix[currPos.x][currPos.y] == "#" {
				isEnd = true
				break
			}
			if matrix[currPos.x][currPos.y] == "." {
				continue
			}
			if matrix[currPos.x][currPos.y] == "[" {
				if !checkIfContains(currLevelStones, currPos) {
					currLevelStones = append(currLevelStones, Box{currPos.x, currPos.y, "["}, Box{currPos.x, currPos.y + 1, "]"})
				}
			}
			if matrix[currPos.x][currPos.y] == "]" {
				if !checkIfContains(currLevelStones, currPos) {
					currLevelStones = append(currLevelStones, Box{currPos.x, currPos.y - 1, "["}, Box{currPos.x, currPos.y, "]"})
				}
			}

		}
		if isEnd {
			return matrix, pos
		}
		// if there aren't any boxes at the current level break the loop
		if len(currLevelStones) == 0 {
			break
		}
		stonesToMove = append(stonesToMove, currLevelStones)

	}
	// don't move anything if we reached # with any box
	if isEnd {
		return matrix, pos
	}
	// move the boxes
	for j := len(stonesToMove) - 1; j >= 0; j-- {
		for _, stone := range stonesToMove[j] {
			matrix[stone.x][stone.y] = "."
			matrix[stone.x+currDir.x][stone.y+currDir.y] = stone.sign
		}
	}
	// move the robot
	currPosRobot = pos
	matrix[currPosRobot.x][currPosRobot.y] = "."
	currPosRobot.x += currDir.x
	currPosRobot.y += currDir.y
	matrix[currPosRobot.x][currPosRobot.y] = "@"

	return matrix, currPosRobot

}

func checkIfContains(lst []Box, pos Point) bool {
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
