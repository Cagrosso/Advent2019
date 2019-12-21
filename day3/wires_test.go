package main

import (
	"testing"
)

func TestClosestIntersection(t *testing.T) {
	type testStruct struct {
		input    string
		expected int
	}

	tests := []testStruct{
		testStruct{
			`R75,D30,R83,U83,L12,D49,R71,U7,L72
			U62,R66,U55,R34,D71,R55,D58,R83`,
			159,
		},
		testStruct{
			`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
			U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			135,
		},
	}

	for i, test := range tests {
		output := closestIntersection(test.input)
		if output != test.expected {
			t.Errorf("Test '%v': ClosestIntersection(%v) expected '%v' got '%v'", i, test.input, test.expected, output)
		}
	}
}

func TestConstructLineSegmentsFromWire(t *testing.T) {
	type testStruct struct {
		input    string
		expected []lineSegment
	}
	tests := []testStruct{
		testStruct{
			"U1,R3",
			[]lineSegment{
				lineSegment{
					point{0, 0},
					point{3, 1},
				},
			},
		},
		testStruct{
			"U1,L3",
			[]lineSegment{
				lineSegment{
					point{0, 0},
					point{-3, 1},
				},
			},
		},
		testStruct{
			"D1,R3,",
			[]lineSegment{
				lineSegment{
					point{0, 0},
					point{3, -1},
				},
			},
		},
		testStruct{
			"D1,L3,",
			[]lineSegment{
				lineSegment{
					point{0, 0},
					point{-3, -1},
				},
			},
		},
		testStruct{
			"D1,L3,U2",
			[]lineSegment{
				lineSegment{
					point{0, 0},
					point{-3, -1},
				},
				lineSegment{
					point{-3, -1},
					point{-3, 1},
				},
			},
		},
	}

	for i, test := range tests {
		output := constructLineSegmentsFromWire(test.input)
		for j, segment := range output {
			if !segment.Equal(test.expected[j]) {
				t.Errorf("Test '%v': ClosestIntersection(%v) expected '%v' got '%v'", i, test.input, test.expected, output)
			}
		}
	}
}
