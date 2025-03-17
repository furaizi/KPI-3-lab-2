package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePostfix_NormalCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Addition", "1 2 +", "3"},
		{"Multiplication", "3 6 *", "18"},
		{"Subtraction", "5 3 -", "2"},
		{"Division", "-8 2 /", "-4"},
		{"Power", "4 3 ^", "64"},
		// (1 + 2) * (3 - 4) / 5
		{"Medium", "1 2 + 3 4 - * 5 /", "-0.6"},
		// 42 / 6 + (70 - 64) * (2 ^ 2) + 52 / 26 - 70 / 5
		{"Hard", "42 6 / 70 64 - 2 2 ^ * + 52 25 / + 70 5 / -", "19"},
		// 5 + (6 * (2 ^ 3 - 4)) / (3 - 1)
		{"Complex", "5 6 2 3 ^ 4 - * 3 1 - / +", "17"},
	}

	for _, test := range tests {
		// capturing the loop variable
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			res, err := CalculatePostfix(test.input)
			require.NoError(t, err, "Unexpected error for input string %q", test.input)
			assert.Equal(t, test.expected, res, "Invalid result for input string %q", test.input)
		})
	}
}

func TestCalculatePostfix_ExceptionalCases(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{"EmptyString", "", emptyStringErrorMessage},
		{"NotEnoughOperands", "9 -", incorrectNumberOfOperandsErrorMessage},
		{"TooManyOperands", "5 1 2 *", incorrectNumberOfOperandsErrorMessage},
		{"NotAnExpression", "abc", unsupportedSymbolErrorMessage},
		{"UnknownOperator", "9 2 &", unsupportedSymbolErrorMessage},
		{"UnknownOperand", "9 b +", unsupportedSymbolErrorMessage},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := CalculatePostfix(test.input)
			require.Error(t, err, "Error expected for input string %q", test.input)
			assert.ErrorContains(t, err, test.expectedErr, "Error message does not match expectation for input string %q", test.input)
		})
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("5 2 +")
	fmt.Println(res)
}
