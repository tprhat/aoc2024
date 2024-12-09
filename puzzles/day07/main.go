package main

import (
	"bufio"
	"fmt"
	"math"
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
	lines := transformInput(name)
	total := 0
	for _, line := range lines {
		operation := strings.Split(line, ": ")
		result, _ := strconv.Atoi(operation[0])
		values := strings.Split(operation[1], " ")
		opList := buildOperations(len(values) - 1)
		for _, op := range opList {
			res, _ := strconv.Atoi(values[0])
			for i := 1; i < len(values); i++ {
				if res > result {
					break
				}
				val, _ := strconv.Atoi(values[i])
				if op[i-1] == 0 {
					res += val
				} else {
					res *= val
				}
			}
			if res == result {
				total += res
				break
			}
		}
	}
	return total
}

// + *
func buildOperations(lenOps int) [][]int {
	op := [][]int{}
	totalOps := int(math.Pow(2, float64(lenOps)))
	for i := range totalOps {
		op1 := []int{}
		for j := range lenOps {
			// divide by 2^j AND with least significant bit
			op1 = append(op1, (i>>j)&1)
		}
		op = append(op, op1)
	}
	return op
}

func part2(name string) int {
	lines := transformInput(name)
	total := 0
	for _, line := range lines {
		operation := strings.Split(line, ": ")
		result, _ := strconv.Atoi(operation[0])
		values := strings.Split(operation[1], " ")
		opList := buildTernaryOperations(len(values) - 1)
		for _, op := range opList {
			res, _ := strconv.Atoi(values[0])
			for i := 1; i < len(values); i++ {
				if res > result {
					break
				}
				val, _ := strconv.Atoi(values[i])
				if op[i-1] == 0 {
					res += val
				} else if op[i-1] == 1 {
					res *= val
				} else {
					res = concat(res, val)
				}
			}
			if res == result {
				total += res
				break
			}
		}
	}
	return total
}

// + * ||
func buildTernaryOperations(lenOps int) [][]int {
	op := [][]int{}
	totalOps := int(math.Pow(3, float64(lenOps)))
	for i := range totalOps {
		op1 := []int{}
		currentNum := i
		for range lenOps {
			op1 = append(op1, currentNum%3)
			currentNum /= 3
		}
		op = append(op, op1)
	}
	return op
}

func concat(a, b int) int {
	tmp := b
	for {
		if tmp == 0 {
			break
		}
		tmp /= 10
		a *= 10
	}
	return a + b
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
