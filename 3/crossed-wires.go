package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// Point struct
type Point struct {
	X, Y int
}

func convertInputToArrays(input string) ([]string, []string) {
	wires := strings.Split(input, "\n")

	return strings.Split(wires[0], ","), strings.Split(wires[1], ",")
}

func mapPointsOfWire(wire []string) map[Point]bool {
	points := make(map[Point]bool)

	current := Point{0, 0}

	points[current] = true

	for _, instruction := range wire {
		direction := string(instruction[0])
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}

		x := current.X
		y := current.Y

		for steps > 0 {
			if direction == "D" {
				y--
			} else if direction == "R" {
				x++
			} else if direction == "U" {
				y++
			} else if direction == "L" {
				x--
			}

			point := Point{x, y}

			points[point] = true

			steps--
		}

		current = Point{x, y}
	}

	return points
}

func manhattanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Abs(x2-x1) + math.Abs(y2-y1)
}

func solve1(wire1 []string, wire2 []string) float64 {
	points1 := mapPointsOfWire(wire1)
	points2 := mapPointsOfWire(wire2)

	minDistance := math.Inf(1)

	for point := range points2 {
		if point.X == 0 && point.Y == 0 {
			continue
		}

		if _, ok := points1[point]; ok {
			if dist := manhattanDistance(0, 0, float64(point.X), float64(point.Y)); dist < minDistance {
				minDistance = dist
			}
		}
	}

	return minDistance
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	wire1, wire2 := convertInputToArrays(string(content))

	minDistance := solve1(wire1, wire2)

	log.Println("Distance of closest crossing point is", minDistance)
}
