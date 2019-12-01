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

func solve1(masses []string) int {
	totalFuelRequirement := 0
	for _, massString := range masses {
		massInt, err := strconv.Atoi(massString)
		if err != nil {
			log.Fatal(err)
		}
		fuel := calculateRequiredFuel(massInt)
		log.Printf("Fuel needed for module with weight %d is %d", massInt, fuel)
		totalFuelRequirement += fuel
	}

	return totalFuelRequirement
}

func solve2(masses []string) int {
	totalFuelRequirement := 0
	for _, massString := range masses {
		massInt, err := strconv.Atoi(massString)
		if err != nil {
			log.Fatal(err)
		}
		fuel := calculateAdjustedRequiredFuel(massInt)
		log.Printf("Adjusted fuel needed for module with weight %d is %d", massInt, fuel)
		totalFuelRequirement += fuel
	}

	return totalFuelRequirement
}

func main() {
	lines, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(lines), "\n")

	result := solve1(input)
	adjustedResult := solve2(input)

	log.Println("Total fuel requirement is", result)
	log.Println("Total adjusted fuel requirement is", adjustedResult)
}
