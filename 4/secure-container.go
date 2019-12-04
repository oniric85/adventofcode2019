package main

import "log"

func isValidPassword(pass int) bool {
	previous := 9
	hasDoubleDigit := false
	for i := 0; i < 6; i++ {
		digit := pass % 10
		pass = pass / 10

		// digits must be non-decrescent from left to right
		if digit > previous {
			return false
		}

		if i != 0 && digit == previous {
			hasDoubleDigit = true
		}

		previous = digit
	}

	return hasDoubleDigit
}

func solve1(start, end int) int {
	validPasswords := 0
	for test := start; test <= end; test++ {
		if isValidPassword(test) {
			validPasswords++
		}
	}

	return validPasswords
}

func main() {
	start := 246515
	end := 739105

	validPasswords := solve1(start, end)

	log.Println("Number of valid passwords is", validPasswords)
}
