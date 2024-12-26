package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// part one was just running all the commands, they are not all sequental so we have to check if
// the variables are initialized first. Finally doing a version of binary arithmetic we come to the solution.
func part1(name string) int {
	inputs, commands := transformInput(name)
	inputs = runCommands(inputs, commands)
	total := 0
	for k, v := range inputs {
		if strings.HasPrefix(k, "z") {
			if v {
				val, _ := strconv.Atoi(k[1:])
				total += int(math.Pow(2, float64(val)))
			}
		}
	}
	return total
}

func pprint(m map[string]bool, comms map[string]Command, reg string, depth int) {
	if depth > 5 {
		return
	}
	exists := false
	for k := range maps.Keys(m) {
		if k == reg {
			exists = true
		}
	}
	if !exists {
		str := ""
		for range depth {
			str += "  "
		}
		fmt.Println(str, reg, comms[reg])
		pprint(m, comms, comms[reg].a, depth+1)
		pprint(m, comms, comms[reg].b, depth+1)
	}
}

// this was a bit more complicated and included doing some reverse engineering to find the values
// that need to be swapped. the puzzle text says that the program is trying to add the number x and y.
// first we need to figure out what bits are wrong, then check what those bits are doing under the hood
// and finally swap the command outputs.
// it turns out tha tall of the correct outputs follow the same pattern of commands
//
//	XOR
//		XOR
//		OR
//			AND
//			AND
//				XOR
//				OR
//					AND
//					AND
//						OR
//						XOR
//
// we could check how the wrong bits operate using the pprint command and see what could be swapped to
// get the output similar to this. the process involves running the code multiple times and looking looking
// at the input for those bits
func part2(name string) string {
	inputs, commands := transformInput(name)

	// varibles to swap for my input
	tmp := commands["gvw"]
	commands["gvw"] = commands["qjb"]
	commands["qjb"] = tmp

	tmp = commands["jgc"]
	commands["jgc"] = commands["z15"]
	commands["z15"] = tmp

	tmp = commands["drg"]
	commands["drg"] = commands["z22"]
	commands["z22"] = tmp

	tmp = commands["jbp"]
	commands["jbp"] = commands["z35"]
	commands["z35"] = tmp

	// used for finding what to swap
	// pprint(inputs, commands, "z14", 0)
	// fmt.Println()
	// pprint(inputs, commands, "z15", 0)
	// fmt.Println()
	// pprint(inputs, commands, "z16", 0)
	// fmt.Println()
	// pprint(inputs, commands, "z18", 0)
	inputs = runCommands(inputs, commands)
	totalX := 0
	totalY := 0
	totalZ := 0
	for k, v := range inputs {
		if strings.HasPrefix(k, "x") {
			if v {
				val, _ := strconv.Atoi(k[1:])
				totalX += int(math.Pow(2, float64(val)))
			}
		}
		if strings.HasPrefix(k, "y") {
			if v {
				val, _ := strconv.Atoi(k[1:])
				totalY += int(math.Pow(2, float64(val)))
			}
		}
		if strings.HasPrefix(k, "z") {
			if v {
				val, _ := strconv.Atoi(k[1:])
				totalZ += int(math.Pow(2, float64(val)))
			}
		}
	}
	// actualZ := totalX + totalY

	// fmt.Println("x=  ", strconv.FormatInt(int64(totalX), 2))
	// fmt.Println("y=  ", strconv.FormatInt(int64(totalY), 2))
	// fmt.Println("z= ", strconv.FormatInt(int64(totalZ), 2))
	// fmt.Println("C= ", strconv.FormatInt(int64(actualZ), 2))
	// fmt.Printf("D=  %046b\n", actualZ^totalZ)
	// 0000000011100000000000110001111000001100000000

	// join the slice and sort it for the final solution
	l := []string{"gvw", "qjb", "jgc", "z15", "drg", "z22", "jbp", "z35"}
	slices.Sort(l)
	return strings.Join(l, ",")
}

func runCommands(in map[string]bool, commands map[string]Command) map[string]bool {
	inputs := make(map[string]bool)
	maps.Copy(inputs, in)
	comms := make(map[string]Command)
	maps.Copy(comms, commands)

	for len(comms) > 0 {
		for k, v := range comms {
			if _, ok := inputs[v.a]; !ok {
				continue
			}
			if _, ok := inputs[v.b]; !ok {
				continue
			}
			val1 := inputs[v.a]
			val2 := inputs[v.b]
			switch v.op {
			case "OR":
				inputs[k] = val1 || val2
			case "AND":
				inputs[k] = val1 && val2
			case "XOR":
				inputs[k] = val1 != val2
			}
			delete(comms, k)
		}
	}
	return inputs
}

type Command struct {
	a, b, op string
}

func transformInput(name string) (map[string]bool, map[string]Command) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	m := make(map[string]bool)
	commands := map[string]Command{}
	cms := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			cms = true
			continue
		}
		if cms {
			str := strings.Split(scanner.Text(), " ")
			commands[str[4]] = Command{str[0], str[2], str[1]}
		} else {
			elems := strings.Split(scanner.Text(), ": ")
			val, _ := strconv.Atoi(elems[1])
			m[elems[0]] = val != 0
		}
	}
	return m, commands
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
