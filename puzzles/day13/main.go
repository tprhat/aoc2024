package main

import (
	"aoc2024/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"gonum.org/v1/gonum/mat"
)

// this task comes down to solving matrix equations
// after changing the input from:
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
// to:
// 94x + 22y = 8400
// 34x + 67y = 5400
// the equations are linear there and there can only be at most one solution
// we solve it using least squares algorithm

type Token struct {
	x, y int
}
type Coords struct {
	x, y int
}

func checkEquality(tmp, res, eq1, eq2 Coords) bool {
	if tmp.x*eq1.x+tmp.y*eq2.x == res.x && tmp.x*eq1.y+tmp.y*eq2.y == res.y {
		return true
	}
	return false
}

func solve(lines []string, args ...bool) Token {
	var eq1, eq2, res Coords
	var tokens Token
	isPart2 := false
	if len(args) > 0 {
		isPart2 = true
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "Button A:") {
			nums := util.GetNumsFromString[int](line)
			eq1.x = nums[0]
			eq1.y = nums[1]
		}
		if strings.HasPrefix(line, "Button B:") {
			nums := util.GetNumsFromString[int](line)
			eq2.x = nums[0]
			eq2.y = nums[1]
		}
		if strings.HasPrefix(line, "Prize:") {
			nums := util.GetNumsFromString[int](line)
			if !isPart2 {
				res.x, res.y = nums[0], nums[1]
			} else {
				res.x, res.y = nums[0]+10000000000000, nums[1]+10000000000000
			}
			// solve matrix
			A := mat.NewDense(2, 2, []float64{float64(eq1.x), float64(eq2.x), float64(eq1.y), float64(eq2.y)})
			b := mat.NewVecDense(2, []float64{float64(res.x), float64(res.y)})
			var x mat.VecDense
			if err := x.SolveVec(A, b); err != nil {
				fmt.Println("no combinations for ", A, b)
				continue
			}

			tmp1 := int(math.Round(x.At(0, 0)))
			tmp2 := int(math.Round(x.At(1, 0)))
			// check if nums are integers
			if checkEquality(Coords{tmp1, tmp2}, res, eq1, eq2) {
				tokens.x += 3 * tmp1
				tokens.y += tmp2
				continue
			}
		}
	}
	return tokens
}
func part1(name string) int {
	lines := transformInput(name)
	tokens := solve(lines)
	return tokens.x + tokens.y
}

func part2(name string) int {
	lines := transformInput(name)
	tokens := solve(lines, true)
	return tokens.x + tokens.y
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
