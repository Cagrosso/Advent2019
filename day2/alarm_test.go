package main

import (
	"fmt"
	"testing"
)

type testStruct struct {
	test     []int
	expected []int
}

func isArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for index, aElement := range a {
		if b[index] != aElement {
			return false
		}
	}

	return true
}

func TestIntCodeRunner(t *testing.T) {
	tests := []testStruct{
		testStruct{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		testStruct{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		testStruct{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		testStruct{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
		testStruct{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}

	for testIndex, test := range tests {
		answer := IntCodeRunner(test.test)
		if !isArraysEqual(answer, test.expected) {
			t.Errorf("Test %v: IntCodeRunner(%v) expected '%v' got '%v'", testIndex, test.test, test.expected, answer)
		}
	}
}

func TestRun1(t *testing.T) {
	test := []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 10, 23, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 35, 9, 39, 2, 39, 13, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 9, 59, 63, 2, 6, 63, 67, 1, 13, 67, 71, 1, 9, 71, 75, 2, 13, 75, 79, 1, 79, 10, 83, 2, 83, 9, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 99, 5, 103, 1, 103, 2, 107, 1, 107, 10, 0, 99, 2, 0, 14, 0}

	result := IntCodeRunner(test)

	fmt.Println(result)
}
