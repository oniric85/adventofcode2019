package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func powInt(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}

	return result
}

func getOpcode(number int) int {
	return number % 100
}

func convertInputToArray(input string) []int {
	tokens := strings.Split(input, ",")

	ints := make([]int, 0, len(tokens))

	for _, line := range tokens {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, n)
	}

	return ints
}

func readIntegerFromInput() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Panic(err)
	}

	return i
}

func getArgument(modes int, pointer int, position int, inOut int, program []int) int {
	mode := (modes % powInt(10, position)) / powInt(10, position-1)

	if mode == 1 || inOut == 1 {
		return program[pointer+position]
	}

	return program[program[pointer+position]]
}

func runProgram(program []int) {
	pos := 0
	for program[pos] != 99 {
		instruction := program[pos]
		opcode := getOpcode(instruction)
		modes := instruction / 100

		if opcode == 1 {
			param1 := getArgument(modes, pos, 1, 0, program)
			param2 := getArgument(modes, pos, 2, 0, program)
			param3 := getArgument(modes, pos, 3, 1, program)
			program[param3] = param1 + param2
			pos += 4
		} else if opcode == 2 {
			param1 := getArgument(modes, pos, 1, 0, program)
			param2 := getArgument(modes, pos, 2, 0, program)
			param3 := getArgument(modes, pos, 3, 1, program)
			program[param3] = param1 * param2
			pos += 4
		} else if opcode == 3 {
			fmt.Println("Provide integer:")
			param1 := getArgument(modes, pos, 1, 1, program)
			program[param1] = readIntegerFromInput()
			pos += 2
		} else if opcode == 4 {
			param1 := getArgument(modes, pos, 1, 0, program)
			fmt.Println("Printing:", param1)
			if param1 == 3 {
				return
			}
			pos += 2
		} else if opcode == 5 {
			param1 := getArgument(modes, pos, 1, 0, program)
			if param1 != 0 {
				param2 := getArgument(modes, pos, 2, 0, program)
				pos = param2
			} else {
				pos += 3
			}
		} else if opcode == 6 {
			param1 := getArgument(modes, pos, 1, 0, program)
			if param1 == 0 {
				param2 := getArgument(modes, pos, 2, 0, program)
				pos = param2
			} else {
				pos += 3
			}
		} else if opcode == 7 {
			param1 := getArgument(modes, pos, 1, 0, program)
			param2 := getArgument(modes, pos, 2, 0, program)
			param3 := getArgument(modes, pos, 3, 1, program)

			if param1 < param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}

			pos += 4
		} else if opcode == 8 {
			param1 := getArgument(modes, pos, 1, 0, program)
			param2 := getArgument(modes, pos, 2, 0, program)
			param3 := getArgument(modes, pos, 3, 1, program)

			if param1 == param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}

			pos += 4
		} else {
			log.Panic("Found opcode ", opcode)
		}
	}
}

func solve1(program []int) {
	runProgram(program)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := convertInputToArray(string(content))

	solve1(input)
}
