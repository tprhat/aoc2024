package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// do the basic blink as per the instructions
// could be done inplace as well
func blink(s []int) []int {
	ns := []int{}
	for _, elem := range s {
		if elem == 0 {
			ns = append(ns, 1)
		} else if len(strconv.Itoa(elem))%2 == 0 {
			e := strconv.Itoa(elem)
			e1 := e[:len(e)/2]
			e2 := e[len(e)/2:]
			e1num, _ := strconv.Atoi(e1)
			e2num, _ := strconv.Atoi(e2)
			ns = append(ns, e1num, e2num)
		} else {
			ns = append(ns, elem*2024)
		}
	}
	return ns
}

type CacheKey struct {
	x int
	n int
}

var cache map[CacheKey]int = make(map[CacheKey]int)

// the problem with this task is that execution time grows exponentially
// so even tho the blink function was good enough for 25 iterations,
// it wasn't sufficient for 75 iterations.
// since the initial elements are independent of each other we can
// calculate them one by one at the required depth and store each step in a cache.
// this way every pair of (x, n) at every level
// first checks the cache and speeds up the process.
func ans(x int, n int) int {
	if n == 0 {
		return 1
	}
	if _, exists := cache[CacheKey{x: x, n: n}]; !exists {
		result := 0
		if x == 0 {
			result = ans(1, n-1)
		} else if len(strconv.Itoa(x))%2 == 0 {
			e := strconv.Itoa(x)
			e1 := e[:len(e)/2]
			e2 := e[len(e)/2:]
			e1num, _ := strconv.Atoi(e1)
			e2num, _ := strconv.Atoi(e2)
			result += ans(e1num, n-1)
			result += ans(e2num, n-1)
		} else {
			result = ans(2024*x, n-1)
		}
		cache[CacheKey{x: x, n: n}] = result
	}
	return cache[CacheKey{x: x, n: n}]

}

func part1(name string) int {
	line := strings.Fields(transformInput(name)[0])
	s := []int{}
	for _, l := range line {
		num, _ := strconv.Atoi(l)
		s = append(s, num)
	}

	for range 25 {
		s = blink(s)
	}

	return len(s)
}

func part2(name string) int {
	line := strings.Fields(transformInput(name)[0])
	s := []int{}
	for _, l := range line {
		num, _ := strconv.Atoi(l)
		s = append(s, num)
	}
	res := 0
	for _, x := range s {
		res += ans(x, 75)
	}
	return res
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
