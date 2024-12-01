package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))

}

func part1(name string) int {
	ids1, ids2 := transformInput(name)
	slices.Sort(ids1)
	slices.Sort(ids2)
	totalDistance := 0
	for i := range ids1 {
		totalDistance += int(math.Abs(float64(ids1[i] - ids2[i])))
	}
	return totalDistance

}

func part2(name string) int {
	ids1, ids2 := transformInput(name)
	// var idsMap map[int]int
	idsMap := make(map[int]int)

	for i := range ids1 {
		for j := range ids2 {
			if ids1[i] == ids2[j] {
				if _, ok := idsMap[ids1[i]]; ok {
					idsMap[ids1[i]]++
				} else {
					idsMap[ids1[i]] = 1
				}
			}
		}
	}

	totalDistance := 0
	for k, v := range idsMap {
		totalDistance += k * v

	}
	return totalDistance
}

func transformInput(name string) ([]int, []int) {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")
	var ids1 []int
	var ids2 []int
	for _, line := range lines {
		num1, err := strconv.Atoi(strings.Fields(line)[0])
		if err != nil {
			fmt.Println(err)
		}
		num2, err := strconv.Atoi(strings.Fields(line)[1])
		if err != nil {
			fmt.Println(err)
		}
		ids1 = append(ids1, num1)
		ids2 = append(ids2, num2)
	}

	return ids1, ids2

}
