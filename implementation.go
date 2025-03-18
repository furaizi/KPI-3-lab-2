package lab2

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var emptyStringErrorMessage = "Error: empty input"
var unsupportedSymbolErrorMessage = "Error: unsupported symbol"
var incorrectNumberOfOperandsErrorMessage = "Error: incorrect number of operands"

func CalculatePostfix(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New(emptyStringErrorMessage)
	}

	tokens := strings.Fields(input)
	stack := []float64{}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
				if len(stack) < 2 {
					return "", errors.New(incorrectNumberOfOperandsErrorMessage)
				}
				b, a := stack[len(stack)-1], stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				var result float64
				switch token {
				case "+":
					result = a + b
				case "-":
					result = a - b
				case "*":
					result = a * b
				case "/":
					if b == 0 {
						return "", errors.New("division by zero")
					}

					if a >= 0 && b > 0 {
						result = float64(int(a) / int(b))
					} else {
						result = a / b
					}
				case "^":
					result = math.Pow(a, b)
				}
				stack = append(stack, result)
			} else {
				return "", errors.New(unsupportedSymbolErrorMessage)
			}
		}
	}

	if len(stack) != 1 {
		return "", errors.New(incorrectNumberOfOperandsErrorMessage)
	}

	return formatResult(stack[0]), nil
}

func formatResult(result float64) string {
	if result == math.Trunc(result) {
		return fmt.Sprintf("%d", int(result))
	}
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", result), "0"), ".")
}
