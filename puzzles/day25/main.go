package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// check all the keys with all locks, pretty easy one for the end
func part1(name string) int {
	keys, locks := transformInput(name)
	total := 0
	for _, key := range keys {
	loop:
		for _, lock := range locks {
			fail := false
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					continue loop
				}
			}
			if !fail {
				total += 1
			}
		}
	}
	return total
}

func part2() string {
	return "That's it for this year! Just hit the button"
}

func parseMatrix(schematics [][]string) []int {
	values := []int{}
	for j := 0; j <= 4; j++ {
		cnt := 0
		for i := 0; i <= 6; i++ {
			if schematics[i][j] == "#" {
				cnt++
			}
		}
		values = append(values, cnt-1)
	}
	return values
}

func transformInput(name string) ([][]int, [][]int) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	schematics := [][]string{}
	keys := [][]int{}
	locks := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			values := parseMatrix(schematics)
			if schematics[0][0] == "#" {
				locks = append(locks, values)
			} else {
				keys = append(keys, values)
			}
			schematics = [][]string{}
			continue
		}
		schematics = append(schematics, strings.Split(scanner.Text(), ""))
	}
	// add the last matrix to the list since it's EOF
	values := parseMatrix(schematics)
	if schematics[0][0] == "#" {
		locks = append(locks, values)
	} else {
		keys = append(keys, values)
	}
	return keys, locks
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
