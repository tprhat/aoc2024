package util

import (
	"math"
	"sort"
	"strconv"
	"unicode"
)

type Number interface {
	~int | ~int64 | ~float64
}

func NextNum[T Number](b string, i int) (T, int, bool) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])) && b[i] != '-' && b[i] != '.'; i++ {
	}
	start := i
	for ; i < len(b) && (unicode.IsDigit(rune(b[i])) || b[i] == '.' || b[i] == '-'); i++ {
	}
	if start == i {
		return 0, i, false
	}
	num, err := strconv.ParseFloat(b[start:i], 64)
	if err != nil {
		return 0, start, false
	}

	return T(num), i, true
}

func GetNumsFromString[T Number](str string) []T {
	var nums []T
	i := 0
	for i < len(str) {
		num, nextI, ok := NextNum[T](str, i)
		if nextI == i {
			break
		}
		if ok {
			nums = append(nums, num)
		}
		i = nextI
	}
	return nums
}

// GetIntsFromStrings extracts all integers from a slice of strings
func GetNumsFromStringSlice[T Number](strs []string) []T {
	var allNums []T
	for _, str := range strs {
		nums := GetNumsFromString[T](str)
		allNums = append(allNums, nums...)
	}
	return allNums
}

type Point struct {
	x, y int
}
type Direction struct {
	dx, dy int
}

var dirs []Direction = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Dijkstra(matrix [][]string, source Point) (map[Point]int, map[Point]Point) {
	dist := make(map[Point]int)
	prev := make(map[Point]Point)
	var Q []Point
	for i := range matrix {
		for j := range matrix[0] {
			if i == 0 && j == 0 {
				dist[source] = 0
				Q = append(Q, source)
				continue
			}
			Q = append(Q, Point{i, j})
			dist[Point{i, j}] = math.MaxInt
		}
	}
	for len(Q) > 0 {
		sort.Slice(Q, func(i, j int) bool {
			return dist[Q[i]] < dist[Q[j]]
		})
		u := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			nx, ny := u.x+dir.dx, u.y+dir.dy
			if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[0]) || matrix[nx][ny] == "#" || !isInQ(Q, nx, ny) || dist[u] == math.MaxInt {
				continue
			}
			v := Point{nx, ny}
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return dist, prev
}
func isInQ(Q []Point, x, y int) bool {
	for _, point := range Q {
		if point.x == x && point.y == y {
			return true
		}
	}
	return false
}
