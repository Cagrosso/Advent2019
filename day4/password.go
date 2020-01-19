package password

import (
	"fmt"
	"strconv"
	"strings"
)

func searchPasswordsMethod1(start, end int) int {
	validPasswords := 0

	for i := start; i <= end; i++ {
		if validatePasswordMethod1(strconv.Itoa(i)) {
			validPasswords++
		}
	}

	return validPasswords
}

func searchPasswordsMethod2(start, end int) int {
	validPasswords := 0

	for i := start; i <= end; i++ {
		if validatePasswordMethod2(strconv.Itoa(i)) {
			validPasswords++
		}
	}

	return validPasswords
}

func validatePasswordMethod1(password string) bool {
	if len(password) != 6 {
		return false
	}

	hasMatchingAdjacent := false

	previousNumber, _ := strconv.Atoi(string(password[0]))

	for i := 1; i < len(password); i++ {
		currentNumber, _ := strconv.Atoi(string(password[i]))

		if currentNumber < previousNumber {
			return false
		}

		if previousNumber == currentNumber {
			hasMatchingAdjacent = true
		}

		previousNumber = currentNumber
	}

	return hasMatchingAdjacent
}

func validatePasswordMethod2(password string) bool {
	if len(password) != 6 {
		return false
	}

	hasMatchingAdjacent := false

	previousNumber, _ := strconv.Atoi(string(password[0]))

	for i := 1; i < len(password); i++ {
		currentNumber, _ := strconv.Atoi(string(password[i]))

		if currentNumber < previousNumber {
			return false
		}

		if previousNumber == currentNumber && !strings.Contains(password, fmt.Sprintf("%v%v%v", currentNumber, currentNumber, currentNumber)) {
			hasMatchingAdjacent = true
		}

		previousNumber = currentNumber
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
