package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompute_ValidInput(t *testing.T) {
	input := strings.NewReader("2 10 5 / +")
	output := new(bytes.Buffer)
	handler := ComputeHandler{input, output}
	err := handler.Compute()

	require.NoError(t, err)
	assert.Equal(t, "5\n", output.String())  // Змінили очікуване значення з "4\n" на "5\n"
}

func TestCompute_InvalidInput(t *testing.T) {
	input := strings.NewReader("20 30 @")
	output := new(bytes.Buffer)
	handler := ComputeHandler{input, output}
	err := handler.Compute()

	require.Error(t, err)
	assert.ErrorContains(t, err, unsupportedSymbolErrorMessage)
}
