package main

import (
	"fmt"
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

func TestTotalFuelCalculation(t *testing.T) {
	tests := []testStruct{
		testStruct{
			1969,
			966,
		},
		testStruct{
			100756,
			50346,
		},
	}

	for _, test := range tests {
		answer := TotalFuelCalculation(test.input)

		if answer != test.output {
			t.Errorf("TotalFuelCalculation(%d) expected '%d' got '%d'", test.input, test.output, answer)
		}
	}
}

func TestRun1(t *testing.T) {
	modules := []int{
		85644,
		52584,
		72349,
		83834,
		56593,
		108492,
		94585,
		97733,
		62732,
		103113,
		133259,
		132647,
		52460,
		51299,
		115749,
		121047,
		69451,
		54737,
		62738,
		116686,
		57293,
		97273,
		128287,
		139440,
		97583,
		130263,
		79307,
		118198,
		82514,
		70679,
		64485,
		119346,
		136281,
		114724,
		73580,
		76314,
		126198,
		97635,
		114655,
		104195,
		99469,
		70251,
		82815,
		79531,
		58135,
		80625,
		73106,
		139806,
		138478,
		136605,
		111472,
		149915,
		95928,
		126905,
		70496,
		147999,
		148501,
		114025,
		75716,
		113473,
		95390,
		104466,
		138715,
		53053,
		79502,
		98601,
		115139,
		122315,
		88402,
		124332,
		140107,
		50912,
		104885,
		142005,
		145938,
		118556,
		101858,
		51142,
		94100,
		99421,
		84544,
		137234,
		126374,
		107333,
		82439,
		125373,
		51212,
		99358,
		82821,
		89913,
		67513,
		136907,
		133707,
		139988,
		96914,
		130672,
		66474,
		120729,
		50131,
		67475,
	}

	total := 0

	for _, module := range modules {
		total += FuelCalculation(module)
	}

	fmt.Printf("The fuel requirement is %v\n", total)
}

func TestRun2(t *testing.T) {
	modules := []int{
		85644,
		52584,
		72349,
		83834,
		56593,
		108492,
		94585,
		97733,
		62732,
		103113,
		133259,
		132647,
		52460,
		51299,
		115749,
		121047,
		69451,
		54737,
		62738,
		116686,
		57293,
		97273,
		128287,
		139440,
		97583,
		130263,
		79307,
		118198,
		82514,
		70679,
		64485,
		119346,
		136281,
		114724,
		73580,
		76314,
		126198,
		97635,
		114655,
		104195,
		99469,
		70251,
		82815,
		79531,
		58135,
		80625,
		73106,
		139806,
		138478,
		136605,
		111472,
		149915,
		95928,
		126905,
		70496,
		147999,
		148501,
		114025,
		75716,
		113473,
		95390,
		104466,
		138715,
		53053,
		79502,
		98601,
		115139,
		122315,
		88402,
		124332,
		140107,
		50912,
		104885,
		142005,
		145938,
		118556,
		101858,
		51142,
		94100,
		99421,
		84544,
		137234,
		126374,
		107333,
		82439,
		125373,
		51212,
		99358,
		82821,
		89913,
		67513,
		136907,
		133707,
		139988,
		96914,
		130672,
		66474,
		120729,
		50131,
		67475,
	}

	total := 0

	for _, module := range modules {
		total += TotalFuelCalculation(module)
	}

	fmt.Printf("The fuel requirement is %v\n", total)
}
