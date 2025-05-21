package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
  inputExpr  = flag.String("e", "", "Expression to evaluate (postfix notation)")
  inputFile  = flag.String("f", "", "Input file with expression")
  outputFile = flag.String("o", "", "Output file for results")
)

func main() {
  flag.Parse()

  if err := validateFlags(); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
  }

  input, err := createInputReader()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
  }

  if c, ok := input.(io.Closer); ok {
     defer func() {
         if err := c.Close(); err != nil {
             fmt.Fprintln(os.Stderr, "error closing input file:", err)
         }
     }()
  }

  output, err := createOutputWriter()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
  }

  if c, ok := output.(io.Closer); ok {
    defer func() {
      if err := c.Close(); err != nil {
        fmt.Fprintln(os.Stderr, "error closing output file")
      }
    }()
  }

  handler := &lab2.ComputeHandler{
    Input:  input,
    Output: output,
  }

  if err := handler.Compute(); err != nil {
    fmt.Fprintln(os.Stderr, "Computation error:", err)
    os.Exit(1)
  }
}

func validateFlags() error {
  if *inputExpr != "" && *inputFile != "" {
    return fmt.Errorf("cannot use both -e and -f")
  }
  if *inputExpr == "" && *inputFile == "" {
    return fmt.Errorf("must provide -e or -f")
  }
  return nil
}

func createInputReader() (io.Reader, error) {
  if *inputExpr != "" {
    return strings.NewReader(*inputExpr), nil
  }

  file, err := os.Open(*inputFile)
  if err != nil {
    return nil, fmt.Errorf("opening input file: %v", err)
  }
  return file, nil
}

func createOutputWriter() (io.Writer, error) {
  if *outputFile == "" {
    return os.Stdout, nil
  }

  file, err := os.Create(*outputFile)
  if err != nil {
    return nil, fmt.Errorf("creating output file: %v", err)
  }
  return file, nil
}