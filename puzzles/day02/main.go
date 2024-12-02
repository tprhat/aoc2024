package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))

}

func part1(name string) int {
	lines := transformInput(name)

	cnt := 0
	for _, line := range lines {
		// separate line by whitespace character
		levelsStr := strings.Fields(line)
		var levels []int
		// convert []string to []int
		for _, level := range levelsStr {
			lvl, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			levels = append(levels, lvl)
		}

		isAsc := false
		if levels[0] < levels[1] {
			isAsc = true
		}
		invalid := false
		for i := 0; i < len(levels)-1; i++ {
			step := levels[i+1] - levels[i]

			if (step > 0) != isAsc {
				invalid = true
				break
			}
			// abs value
			if step < 0 {
				step *= -1
			}

			if step < 1 || step > 3 {
				invalid = true
				break
			}

		}
		if !invalid {
			cnt++
		}
	}
	return cnt
}

func part2(name string) int {
	lines := transformInput(name)

	cnt := 0
	for _, line := range lines {
		levelsStr := strings.Fields(line)
		var levels []int
		for _, level := range levelsStr {
			lvl, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			levels = append(levels, lvl)
		}

		for i := range levels {

			lvls := make([]int, len(levels)-1)

			// copy all elements besides one and make the same check as in part 1
			copy(lvls, levels[:i])
			copy(lvls[i:], levels[i+1:])

			isAsc := false
			if lvls[0] < lvls[1] {
				isAsc = true
			}
			invalid := false
			for i := 0; i < len(lvls)-1; i++ {
				step := lvls[i+1] - lvls[i]

				if (step > 0) != isAsc {
					invalid = true
					break
				}
				// abs value
				if step < 0 {
					step *= -1
				}

				if step < 1 || step > 3 {
					invalid = true
					break
				}

			}
			if !invalid {
				cnt++
				break
			}

		}
	}
	return cnt

}
func transformInput(name string) []string {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\r\n")

	return lines

}
