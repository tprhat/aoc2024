package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// this is a dp problem, we can track a number of ways a string can form from start to finish
// dp[0] being the empty string

// example:
//
// patterns:
// r, wr, b, g, bwu, rb, gb, br
// design:
// brwrr
// solution:
// b r wr r
// br wr r

func waysToCreateLine(line string, patterns []string) int {
	n := len(line)
	dp := make([]int, n+1) // dp[i] = number of ways to create line[0:i]
	dp[0] = 1              // Base case : 1 way to form empty string

	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			pLen := len(pattern)
			if i >= pLen && line[i-pLen:i] == pattern {
				// add the dp[i-pLen] because we have to take into
				// account the length of the pattern we are currently using
				dp[i] += dp[i-pLen]
			}
		}
	}
	return dp[n] // number of ways to form entire string
}

func part1(name string) int {
	lines := transformInput(name)
	patterns := strings.Split(lines[0], ", ")
	lines = lines[2:]
	total := 0
	for _, line := range lines {
		if waysToCreateLine(line, patterns) > 0 {
			total += 1
		}
	}
	return total
}

func part2(name string) int {
	lines := transformInput(name)
	patterns := strings.Split(lines[0], ", ")
	lines = lines[2:]
	total := 0
	for _, line := range lines {
		total += waysToCreateLine(line, patterns)
	}
	return total
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
