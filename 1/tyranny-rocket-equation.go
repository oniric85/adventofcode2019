package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func solve1(input string) int {
	masses := strings.Split(input, "\n")

	totalFuelRequirement := 0
	for _, massString := range masses {
		massInt, err := strconv.Atoi(massString)
		if err != nil {
			log.Fatal(err)
		}
		fuel := massInt/3 - 2
		log.Printf("Fuel needed for module with weight %d is %d", massInt, fuel)
		totalFuelRequirement += fuel
	}

	return totalFuelRequirement
}

func main() {
	lines, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := solve1(string(lines))

	log.Println("Total fuel requirement is", result)
}
