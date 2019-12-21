package main

import (
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
