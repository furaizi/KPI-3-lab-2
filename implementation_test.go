package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePostfix_Addition(t *testing.T) {
	res, err := CalculatePostfix("1 2 +")
	require.Nil(t, err)
	assert.Equal(t, "3", res)
}

func TestCalculatePostfix_Multiplication(t *testing.T) {
	res, err := CalculatePostfix("3 6 *")
	require.Nil(t, err)
	assert.Equal(t, "18", res)
}

func TestCalculatePostfix_Subtraction(t *testing.T) {
	res, err := CalculatePostfix("5 3 -")
	require.Nil(t, err)
	assert.Equal(t, "2", res)
}

func TestCalculatePostfix_Division(t *testing.T) {
	res, err := CalculatePostfix("-8 2 /")
	require.Nil(t, err)
	assert.Equal(t, "-4", res)
}

func TestCalculatePostfix_Power(t *testing.T) {
	res, err := CalculatePostfix("4 3 ^")
	require.Nil(t, err)
	assert.Equal(t, "64", res)
}

func TestCalculatePostfix_Medium(t *testing.T) {
	// (1 + 2) * (3 - 4) / 5
	res, err := CalculatePostfix("1 2 + 3 4 - * 5 /")
	require.Nil(t, err)
	assert.Equal(t, "-0.6", res)
}

func TestCalculatePostfix_Hard(t *testing.T) {
	// 42 / 6 + (70 - 64) * (2 ^ 2) + 52 / 26 - 70 / 5
	res, err := CalculatePostfix("42 6 / 70 64 - 2 2 ^ * + 52 25 / + 70 5 / -")
	require.Nil(t, err)
	assert.Equal(t, "19", res)
}

func TestCalculatePostfix_Complex(t *testing.T) {
	// 5 + (6 * (2 ^ 3 - 4)) / (3 - 1)
	res, err := CalculatePostfix("5 6 2 3 ^ 4 - * 3 1 - / +")
	require.Nil(t, err)
	assert.Equal(t, "17", res)
}

func TestCalculatePostfix_EmptyString(t *testing.T) {
	_, err := CalculatePostfix("")
	if assert.NotNil(t, err) {
		assert.ErrorContains(t, err, emptyStringErrorMessage)
	}
}

func TestCalculatePostfix_NotEnoughOperands(t *testing.T) {
	_, err := CalculatePostfix("9 -")
	if assert.NotNil(t, err) {
		assert.ErrorContains(t, err, incorrectNumberOfOperandsErrorMessage)
	}
}

func TestCalculatePostfix_TooManyOperands(t *testing.T) {
	_, err := CalculatePostfix("5 1 2 *")
	if assert.NotNil(t, err) {
		assert.ErrorContains(t, err, incorrectNumberOfOperandsErrorMessage)
	}
}

func TestCalculatePostfix_UnknownOperator(t *testing.T) {
	_, err := CalculatePostfix("9 2 &")
	if assert.NotNil(t, err) {
		assert.ErrorContains(t, err, unsupportedSymbolErrorMessage)
	}
}

func TestCalculatePostfix_UnknownOperand(t *testing.T) {
	_, err := CalculatePostfix("9 b +")
	if assert.NotNil(t, err) {
		assert.ErrorContains(t, err, unsupportedSymbolErrorMessage)
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
