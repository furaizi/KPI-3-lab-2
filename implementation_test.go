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

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
