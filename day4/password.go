package password

import (
	"strconv"
)

func searchPasswords(start, end int) int {
	validPasswords := 0

	for i := start; i <= end; i++ {
		if validatePassword(strconv.Itoa(i)) {
			validPasswords++
		}
	}

	return validPasswords
}

func validatePassword(passwordString string) bool {
	checkPassword := passwordString
	if len(checkPassword) != 6 {
		checkPassword = padPassword(checkPassword)
	}

	hasMatchingAdjacent := false

	for i := 0; i < len(checkPassword); i++ {
		currentNumber, err := strconv.Atoi(string(checkPassword[i]))
		if err != nil {
			return false
		}

		for j := i + 1; j < len(checkPassword); j++ {
			nextNumber, err := strconv.Atoi(string(checkPassword[j]))
			if err != nil {
				return false
			}

			if j == i+1 {
				if currentNumber == nextNumber {
					hasMatchingAdjacent = true
				}
			}

			if nextNumber < currentNumber {
				return false
			}
		}

	}

	return hasMatchingAdjacent
}

func padPassword(shortPassword string) string {
	paddedPassword := shortPassword

	for i := len(shortPassword); i < 6; i++ {
		paddedPassword = "0" + paddedPassword
	}

	return paddedPassword
}
