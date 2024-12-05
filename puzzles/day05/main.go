package main

import (
	"fmt"
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
	rules, pagesUpdate := transformInput(name)
	total := 0
	for _, pages := range pagesUpdate {
		var pagesPassed []int
		goodPage := true
		for _, page := range strings.Split(pages, ",") {
			pg, _ := strconv.Atoi(page)
			good := true
			for _, pagePassed := range pagesPassed {
				if slices.Contains(rules[pg], pagePassed) {
					good = false
					break
				}
			}
			pagesPassed = append(pagesPassed, pg)
			if !good {
				goodPage = false
				break
			}
		}
		if goodPage {
			currPages := strings.Split(pages, ",")
			val, _ := strconv.Atoi(currPages[len(currPages)/2])
			total += val
		}
	}
	return total
}

func part2(name string) int {
	rules, pagesUpdate := transformInput(name)
	total := 0
	var badPages []string
	for _, pages := range pagesUpdate {
		var pagesPassed []int
		for _, page := range strings.Split(pages, ",") {
			pg, _ := strconv.Atoi(page)
			good := true
			for _, pagePassed := range pagesPassed {
				if slices.Contains(rules[pg], pagePassed) {
					good = false
					break
				}
			}
			pagesPassed = append(pagesPassed, pg)
			if !good {
				badPages = append(badPages, pages)
				break
			}
		}
	}
	for _, pagesStr := range badPages {
		pages := strings.Split(pagesStr, ",")
		pagesSorted := make([]int, len(pages))

		for i, page := range pages {
			pageInt, _ := strconv.Atoi(page)
			cnt := 0
			for _, pg := range pagesSorted {
				if pg == 0 {
					break
				}
				if slices.Contains(rules[pageInt], pg) {
					cnt++
				}
			}
			// shift cnt number of the most right elements
			for j := i - 1; j > i-1-cnt; j-- {
				pagesSorted[j+1] = pagesSorted[j]
			}
			// put the element in the correct position
			pagesSorted[i-cnt] = pageInt

			if i == len(pages)-1 {
				total += pagesSorted[len(pagesSorted)/2]
			}
		}
	}
	return total
}

func transformInput(name string) (map[int][]int, []string) {
	inputFile, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputFile), "\n")
	var pages []string
	rules := make(map[int][]int)

	isRules := true
	for _, line := range lines {
		if line == "" {
			isRules = false
			continue
		}
		if isRules {
			rule := strings.Split(line, "|")
			key, _ := strconv.Atoi(rule[0])
			val, _ := strconv.Atoi(rule[1])

			rules[key] = append(rules[key], val)
		} else {
			pages = append(pages, line)
		}

	}
	return rules, pages
}
