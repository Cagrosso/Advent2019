package password

import (
	"fmt"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	type testStruct struct {
		input  string
		output bool
	}

	tests := []testStruct{
		testStruct{
			"000000",
			true,
		},
		testStruct{
			"112345",
			true,
		},
		testStruct{
			"654322",
			false,
		},
		testStruct{
			"135799",
			true,
		},
		testStruct{
			"0",
			false,
		},
		testStruct{
			"123456",
			false,
		},
	}

	for i, test := range tests {
		result := validatePassword(test.input)

		if result != test.output {
			t.Errorf("Test %v: validatePassword(%v) expected '%v' got '%v'", i, test.input, test.output, result)
		}

		fmt.Println()
	}
}

func TestPadPassword(t *testing.T) {
	test := "123"
	expectedValue := "000123"
	returnValue := padPassword(test)

	if returnValue != expectedValue {
		t.Errorf("padPassword(%v) expected '%v' got '%v'", test, expectedValue, returnValue)
	}
}

func TestSearchPasswords(t *testing.T) {
	type passwordRange struct {
		start int
		end   int
	}
	type testStruct struct {
		input  passwordRange
		output int
	}

	tests := []testStruct{
		testStruct{
			passwordRange{
				0,
				10,
			},
			0,
		},
	}

	for _, test := range tests {
		result := searchPasswords(test.input.start, test.input.end)
		if result != test.output {
			t.Errorf("searchPassword(%v) expected '%v' got '%v'", test.input, test.output, result)
		}
	}
}

func TestRun1(t *testing.T) {
	fmt.Println(searchPasswords(124075, 580769))
}
