package main

import "log"

func isValidPasswordForFirstPart(pass int) bool {
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

func isValidPasswordForSecondPart(pass int) bool {
	previous := 9
	hasRepeatedDoubleDigits := false
	sameDigitInARow := 1
	for i := 0; i < 6; i++ {
		digit := pass % 10
		pass = pass / 10

		// digits must be non-decrescent from left to right
		if digit > previous {
			return false
		}

		if i != 0 {
			if digit == previous {
				sameDigitInARow++
			} else {
				if sameDigitInARow == 2 {
					hasRepeatedDoubleDigits = true
				}
				sameDigitInARow = 1
			}
		}

		previous = digit
	}

	return hasRepeatedDoubleDigits || sameDigitInARow == 2
}

func solve1(start, end int) int {
	validPasswords := 0
	for test := start; test <= end; test++ {
		if isValidPasswordForFirstPart(test) {
			validPasswords++
		}
	}

	return validPasswords
}

func solve2(start, end int) int {
	validPasswords := 0
	for test := start; test <= end; test++ {
		if isValidPasswordForSecondPart(test) {
			validPasswords++
		}
	}

	return validPasswords
}

func main() {
	start := 246515
	end := 739105

	validPasswordsPart1 := solve1(start, end)
	validPasswordsPart2 := solve2(start, end)

	log.Println("Number of valid passwords for part 1 is", validPasswordsPart1)
	log.Println("Number of valid passwords for part 2 is", validPasswordsPart2)
}
