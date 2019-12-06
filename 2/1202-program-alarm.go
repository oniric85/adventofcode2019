package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func runProgram(program []int, noun int, verb int) int {
	program[1] = noun
	program[2] = verb

	pos := 0
	for program[pos] != 99 {
		pos1 := program[pos+1]
		pos2 := program[pos+2]
		dest := program[pos+3]

		if program[pos] == 1 {
			program[dest] = program[pos1] + program[pos2]
		} else if program[pos] == 2 {
			program[dest] = program[pos1] * program[pos2]
		} else {
			log.Panic("Found opcode ", program[pos])
		}

		pos += 4
	}

	return program[0]
}

func solve1(program []int) int {
	// set initial state 1202
	return runProgram(program, 12, 2)
}

func solve2(program []int, wanted int) (int, error) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			copied := make([]int, len(program))
			copy(copied, program)
			result := runProgram(copied, i, j)
			if result == wanted {
				return 100*i + j, nil
			}
		}
	}

	return 0, fmt.Errorf("Could not find input to get %d", wanted)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := convertInputToArray(string(content))
	input2 := make([]int, len(input))
	copy(input2, input)

	result := solve1(input)

	wanted := 19690720
	result2, err := solve2(input2, wanted)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Value at position 0 after program execution is", result)
	log.Printf("Input needed to obtain %d is %d", wanted, result2)
}
