package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))

}

func part1(name string) int {
	lines := transformInput(name)

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	total := 0
	for _, line := range lines {
		submatches := r.FindAllStringSubmatch(line, -1)
		for _, submatch := range submatches {
			num1, _ := strconv.Atoi(submatch[1])
			num2, _ := strconv.Atoi(submatch[2])
			total += num1 * num2
		}
	}
	return total
}

func part2(name string) int {
	lines := transformInput(name)
	r, _ := regexp.Compile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	total := 0
	canMul := true
	for _, line := range lines {
		submatches := r.FindAllStringSubmatch(line, -1)
		for _, submatch := range submatches {
			if submatch[0] == "do()" {
				canMul = true
			} else if submatch[0] == "don't()" {
				canMul = false
			} else {
				if canMul {
					num1, _ := strconv.Atoi(submatch[2])
					num2, _ := strconv.Atoi(submatch[3])
					total += num1 * num2
				}
			}
		}
	}
	return total
}

func transformInput(name string) []string {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")
	return lines
}
