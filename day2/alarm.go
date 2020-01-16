package main

const addOpcode = 1
const multOpcode = 2
const finishedOpcode = 99

const operatorOffset = 0
const operandOneOffset = 1
const operandTwoOffset = 2
const solutionOffset = 3

type program struct {
	operator        int
	operandOneIndex int
	operandTwoIndex int
	solutionIndex   int
}

// IntCodeRunner executes a given IntCode
func IntCodeRunner(code []int) []int {
	var p program

	for index, value := range code {
		codeIndex := index % 4

		switch codeIndex {
		case operatorOffset:
			p.operator = value
		case operandOneOffset:
			p.operandOneIndex = value
		case operandTwoOffset:
			p.operandTwoIndex = value
		case solutionOffset:
			p.solutionIndex = value

			switch p.operator {
			case addOpcode:
				code[p.solutionIndex] = code[p.operandOneIndex] + code[p.operandTwoIndex]
			case multOpcode:
				code[p.solutionIndex] = code[p.operandOneIndex] * code[p.operandTwoIndex]
			case finishedOpcode:
				break
			default:
				break
			}
		default:
			break
		}
	}

	return code
}
