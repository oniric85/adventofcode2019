package main

import (
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

func solve1(program []int) int {
	// set initial state 1202
	program[1] = 12
	program[2] = 2

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
			log.Panic("Found opcode", program[pos])
		}

		pos += 4
	}

	return program[0]
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := convertInputToArray(string(content))

	result := solve1(input)

	log.Println("Value at position 0 after program execution is", result)
}
