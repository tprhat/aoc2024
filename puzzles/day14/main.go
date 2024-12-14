package main

import (
	"aoc2024/util"
	"bufio"
	"fmt"
	"os"
	"time"
)

type Point struct {
	x, y int
}

func part1(name string) int {
	lines := transformInput(name)
	height := 101
	width := 103
	// test input heigth
	// height := 11
	// width := 7
	// count the number elements at the x,y coords after 100 iterations for all lines
	locs := make(map[Point]int)
	for _, line := range lines {
		nums := util.GetNumsFromString[int](line)
		currx, curry := nums[0], nums[1]
		for range 100 {
			currx = (currx + nums[2])
			// if an element gets off map it loops on the other side
			if currx < 0 {
				currx += height
			}
			if currx >= height {
				currx -= height
			}
			curry = (curry + nums[3])
			if curry < 0 {
				curry += width
			}
			if curry >= width {
				curry -= width
			}
		}
		locs[Point{currx, curry}] += 1
	}
	// calculate the elements per quadrants with the middle being omitted
	topleft, bottomleft, topright, bottomright := 0, 0, 0, 0
	for l := range locs {
		if l.x < (height-1)/2 && l.y < (width-1)/2 {
			topleft += locs[l]
		}
		if l.x > (height-1)/2 && l.y < (width-1)/2 {
			bottomleft += locs[l]
		}

		if l.x < (height-1)/2 && l.y > (width-1)/2 {
			topright += locs[l]
		}

		if l.x > (height-1)/2 && l.y > (width-1)/2 {
			bottomright += locs[l]
		}
	}
	return topleft * topright * bottomleft * bottomright
}

// here we have to visually find a picture of a Christmas tree
func part2(name string) int {
	lines := transformInput(name)
	height := 101
	width := 103
	ls := [][]int{}
	for _, line := range lines {
		nums := util.GetNumsFromString[int](line)
		ls = append(ls, nums)
	}
	i := 0
	for {
		i++
		nLineElems := make(map[int]int)
		for i, l := range ls {
			currx, curry := l[0], l[1]
			currx = (currx + l[2])
			if currx < 0 {
				currx += height
			}
			if currx >= height {
				currx -= height
			}
			curry = (curry + l[3])
			if curry < 0 {
				curry += width
			}
			if curry >= width {
				curry -= width
			}
			ls[i] = []int{currx, curry, l[2], l[3]}
			// count the number of elements in a single row
			// since I saw that the rows in images that were converging
			// had more than 20 elems in them
			nLineElems[currx] += 1
		}
		// this was used to find the picture
		// uncomment if necessary
		// print := false
		// for _, v := range nLineElems {
		// 	if v > 20 {
		// 		print = true
		// 		break
		// 	}
		// }
		// my result was at i == 7709 so this is here
		if i == 7709 { // if print {
			for i := 0; i < height; i++ {
				fmt.Println()
				for j := 0; j < width; j++ {
					if isIn(j, i, ls) {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
			}
			fmt.Println()
			fmt.Println(i)
			break
		}
	}
	return i
}
func isIn(i, j int, ls [][]int) bool {
	for _, l := range ls {
		if l[0] == i && l[1] == j {
			return true
		}
	}
	return false
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
