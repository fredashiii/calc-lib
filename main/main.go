package main

import (
	"io"
	"os"
	"strconv"

	"calc-lib/calc"
)

type Handler struct {
	Writer     io.Writer
	Calculator calc.Calculator
}

func (h Handler) Handle(args []string) {
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	result := h.Calculator.Calculate(a, b)
	h.Writer.Write([]byte(strconv.Itoa(result)))
}

func main() {
	h := Handler{
		Writer:     os.Stdout,
		Calculator: calc.Addition{},
	}
	h.Handle(os.Args[1:])
}
