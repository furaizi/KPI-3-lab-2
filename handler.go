package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
    Input  io.Reader
    Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
    exprBytes, err := io.ReadAll(ch.Input)
    if err != nil {
        return err
    }

    expr := string(exprBytes)
    result, err := CalculatePostfix(expr)
    if err != nil {
        return err
    }

    _, err = fmt.Fprintln(ch.Output, result)
    return err
}