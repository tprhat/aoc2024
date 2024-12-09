package main

import (
	"fmt"
	"os"
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
	line := transformInput(name)[0]

	data := []int{}
	cnt := 0
	for i, rune := range line {

		if i%2 == 0 {
			for range int(rune - '0') {
				data = append(data, cnt)
			}
			cnt++
		} else {
			for range int(rune - '0') {
				data = append(data, -1)
			}
		}
	}
	// the idea was to go from the front and back at the same time
	// if it the loop comes to a -1 (empty) it goes from the back
	// and multiplies and sums the values while lowering the lendata
	// otherwise just increase i
	total := 0
	lendata := len(data)
	i := 0
	for i < lendata {
		if data[i] == -1 {
			for {
				if data[lendata-1] == -1 {
					lendata -= 1
					continue
				}
				total += i * data[lendata-1]
				lendata -= 1
				break
			}

		} else {
			total += i * data[i]
		}
		i++
	}
	return total
}

func part2(name string) int {
	line := transformInput(name)[0]

	emptyStarting := [][]int{}      // [startingidx len]
	starting := make(map[int][]int) // val: [startingidx len]
	cnt := 0
	totalLen := 0

	for i, rune := range line {
		ln := int(rune - '0')
		if i%2 == 0 {
			starting[cnt] = []int{totalLen, ln}
			totalLen += ln
			cnt++
		} else {
			emptyStarting = append(emptyStarting, []int{totalLen, ln})
			totalLen += ln
		}
	}
	keyMax := cnt - 1
	// go from the back and check if there is space somewhere before current location
	for j := keyMax; j >= 0; j-- {
		valLen := starting[j][1]
		for i := range emptyStarting {
			// if new id > current current id break
			// we don't want to move data to the right, only to the left
			if emptyStarting[i][0] > starting[j][0] {
				break
			}
			// this way I can keep the starting postion of empty (free) space
			// and the length from that id,
			if emptyStarting[i][1] >= valLen {
				starting[j] = []int{emptyStarting[i][0], starting[j][1]}
				if emptyStarting[i][1] > valLen {
					emptyStarting[i][1] -= valLen
					emptyStarting[i][0] += valLen
				} else {
					// if there is no more at that postion, remove it from the list
					emptyStarting = append(emptyStarting[:i], emptyStarting[i+1:]...)
				}
				break
			}
		}
	}

	total := 0
	for k, v := range starting {
		for i := v[0]; i < v[0]+v[1]; i++ {
			total += k * i
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
