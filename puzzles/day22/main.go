package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func prune(num int) int {
	return num % 16777216
}
func mix(num1, num2 int) int {
	return num1 ^ num2
}

// run the sim and add all values
func part1(name string) int {
	nums := transformInput(name)
	total := 0
	for _, num := range nums {
		for range 2000 {
			num = prune(mix(num, num<<6))
			num = mix(num, num>>5)
			num = prune(mix(num, num<<11))
		}
		total += num
	}
	return total
}
func get_key(l []int) string {
	return fmt.Sprintf("%v", l)
}

// the idea was to have a "global" map that tracks the sum of prices for each change
// and a local map that tracks the max prices for each change in that iteration.
// the max price is just the highest value in the global map
func part2(name string) int {
	nums := transformInput(name)
	totals := make(map[string]int)
	for _, num := range nums {
		visited := make(map[string]bool)
		currChange := make([]int, 4)
		var last int = num % 10
		i := 0
		for range 2000 {
			num = prune(mix(num, num<<6))
			num = mix(num, num>>5)
			num = prune(mix(num, num<<11))

			currChange = currChange[1:]
			currChange = append(currChange, num%10-last)
			last = num % 10
			// skip first 3 since we don't have last 4 chages without it
			if i < 3 {
				i += 1
				continue
			}
			// skip if seen before
			if visited[get_key(currChange)] {
				continue
			}
			visited[get_key(currChange)] = true
			totals[get_key(currChange)] += num % 10
		}
	}
	maxPrice := 0
	for _, v := range totals {
		if v > maxPrice {
			maxPrice = v
		}
	}
	return maxPrice
}

func transformInput(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	nums := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
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
