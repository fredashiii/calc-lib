package calc

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}

func (add Addition) Calculate(a, b int) int {
	return a + b
}

func (sub Subtraction) Calculate(a, b int) int {
	return a - b
}

func (sub Multiplication) Calculate(a, b int) int {
	return a * b
}

func (div Division) Calculate(a, b int) int {
	return a / b
}
