package calc

import (
	"testing"
)

func TestCalculateEasyAdd(t *testing.T) {
	got := Addition{}.Calculate(1, 2)
	want := 3
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateMediumAdd(t *testing.T) {
	got := Addition{}.Calculate(100, 4)
	want := 104
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateEasySubtract(t *testing.T) {
	got := Subtraction{}.Calculate(1, 2)
	want := -1
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateEasyMultiply(t *testing.T) {
	got := Multiplication{}.Calculate(1, 2)
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateEasyDivide(t *testing.T) {
	got := Division{}.Calculate(4, 2)
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
