package calc

import (
	"testing"
)

func TestCalculateEasy(t *testing.T) {
	got := Addition{}.Calculate(1, 2)
	want := 3
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalculateMedium(t *testing.T) {
	got := Addition{}.Calculate(100, 4)
	want := 104
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
