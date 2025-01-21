package calc

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct{}

func (add Addition) Calculate(a, b int) int {
	return a + b
}
