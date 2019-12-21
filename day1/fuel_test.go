package main

import (
	"testing"
)

type testStruct struct {
	input  int
	output int
}

func TestFuelCalculation(t *testing.T) {
	tests := []testStruct{
		testStruct{
			12,
			2,
		},
		testStruct{
			14,
			2,
		},
		testStruct{
			1969,
			654,
		},
		testStruct{
			100756,
			33583,
		},
	}

	for _, test := range tests {
		answer := FuelCalculation(test.input)

		if answer != test.output {
			t.Errorf("FuelCalculation(%d) expected '%d' got '%d'", test.input, test.output, answer)
		}
	}
}
