package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateRequiredFuel(mass int) int {
	return mass/3 - 2
}

func calculateAdjustedRequiredFuel(mass int) int {
	fuel := calculateRequiredFuel(mass)

	if fuel <= 0 {
		return 0
	}

	return fuel + calculateAdjustedRequiredFuel(fuel)
}

func solve1(masses []int) int {
	totalFuelRequirement := 0
	for _, mass := range masses {
		fuel := calculateRequiredFuel(mass)
		log.Printf("Fuel needed for module with weight %d is %d", mass, fuel)
		totalFuelRequirement += fuel
	}

	return totalFuelRequirement
}

func solve2(masses []int) int {
	totalFuelRequirement := 0
	for _, mass := range masses {
		fuel := calculateAdjustedRequiredFuel(mass)
		log.Printf("Adjusted fuel needed for module with weight %d is %d", mass, fuel)
		totalFuelRequirement += fuel
	}

	return totalFuelRequirement
}

func convertInputToArray(input string) []int {
	lines := strings.Split(input, "\n")

	ints := make([]int, 0, len(lines))

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, n)
	}

	return ints
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := convertInputToArray(string(content))

	result := solve1(input)
	adjustedResult := solve2(input)

	log.Println("Total fuel requirement is", result)
	log.Println("Total adjusted fuel requirement is", adjustedResult)
}
