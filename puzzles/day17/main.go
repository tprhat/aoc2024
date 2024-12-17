package main

import (
	"aoc2024/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Register struct {
	A, B, C int
}

var registers Register

func combo_op(o int) int {
	if o == 4 {
		return registers.A
	}
	if o == 5 {
		return registers.B
	}
	if o == 6 {
		return registers.C
	}
	return o
}
func adv(o int) {
	// in-place
	registers.A >>= combo_op(o)
}

func bxl(o int) {
	registers.B ^= o
}

func bst(o int) {
	registers.B = combo_op(o) % 8
}

func jnz(o int) int {
	if registers.A != 0 {
		return o
	}
	return -1
}
func bxc() {
	registers.B ^= registers.C
}
func bdv(o int) {
	// in-place
	registers.B = registers.A >> combo_op(o)
}
func cdv(o int) {
	// in-place
	registers.C = registers.A >> combo_op(o)
}

// A becomes A
func run(registerA int, program []int) []int {
	var out []int
	i := 0
	registers.A = registerA
	registers.B = 0
	registers.C = 0
	for i < len(program) {
		opcode, operator := program[i], program[i+1]
		if opcode == 0 {
			adv(operator)
		}
		if opcode == 1 {
			bxl(operator)
		}
		if opcode == 2 {
			bst(operator)
		}
		if opcode == 3 {
			val := jnz(operator)
			if val != -1 {
				i = val
				continue
			}
		}
		if opcode == 4 {
			bxc()
		}
		if opcode == 5 {
			out = append(out, (combo_op(operator) % 8))
		}
		if opcode == 6 {
			bdv(operator)
		}
		if opcode == 7 {
			cdv(operator)
		}
		i += 2
	}
	return out
}

func part1(name string) string {
	reg, program := transformInput(name)
	registers = reg
	out := run(reg.A, program)
	var fin string
	for i, s := range out {
		ss := strconv.Itoa(s)
		fin += ss
		if i != len(out)-1 {
			fin += ","
		}
	}
	return fin
}

// program = [2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0]
// 2,4 : B = A % 8
// 1,5 : B = B ^ 5
// 7,5 : C = A // (2 ** B)
// 1,6 : B = B ^ 6
// 4,2 : B = B ^ C
// 5,5 : output B % 8
// 0,3 : A = A // (2 ** 3) -- this is the most important instruction for part 2
// 3,0 : move i to 0

// since A always becomes A // (2 ** 3) or A >> 3, the 3 rightmost bits get flushed.
// knowing this we can go from the back of the program and find all the values that are good for the last bit,
// then take the possible values and have them as an input for the last 2 bits and so on...
// the updates for A are :: A = (A_prev << 3) + block :: where block is from 0 to 7 (the bits that were flushed).
func part2(name string) int {
	reg, program := transformInput(name)
	registers = reg

	// first we initialize the possible values for A 0 through 7
	regAPossibilities := make(map[int]bool)
	for i := 0; i < 8; i++ {
		regAPossibilities[i] = true
	}
	for p := 2; p < len(program)+1; p++ {
		// we look at the bits from the back
		goal := program[len(program)-p:]

		nextPossibilites := make(map[int]bool)

		// iterate over possible regA values
		for regABase := range regAPossibilities {
			// block adds the flushed bits
			for block := 0; block < 8; block++ {
				regA := (regABase << 3) + block

				rn := run(regA, program)

				if compareSlices(rn, goal) {
					nextPossibilites[regA] = true
				}
			}
		}
		regAPossibilities = nextPossibilites
	}
	// find the minimul value for the repetition
	minRa := math.MaxInt
	for ra := range regAPossibilities {
		if ra < minRa {
			minRa = ra
		}
	}
	return minRa
}

func compareSlices(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func transformInput(name string) (Register, []int) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	var registers Register
	program := []int{}
	scanner := bufio.NewScanner(file)
	isProgram := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			isProgram = true
			continue
		}
		if isProgram {
			program = util.GetNumsFromString[int](scanner.Text())
		} else {
			lines := strings.Split(scanner.Text(), ": ")
			val, _ := strconv.Atoi(lines[1])
			if strings.HasSuffix(lines[0], "A") {
				registers.A = val
			}
			if strings.HasSuffix(lines[0], "B") {
				registers.B = val
			}
			if strings.HasSuffix(lines[0], "C") {
				registers.C = val
			}
		}
	}
	return registers, program
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
