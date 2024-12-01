package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))

}

func part1(name string) int {
	return 0
}

func part2(name string) int {
	return 0
}

func transformInput(name string) []string {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")

	return lines

}
