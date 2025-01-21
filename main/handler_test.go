package main

import (
	"io"
	"os"
	"testing"

	"calc-lib/calc"
)

func TestHandler_Handle_Easy(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	h := Handler{
		Writer:     os.Stdout,
		Calculator: calc.Addition{},
	}
	nums := []string{"1", "1"}
	h.Handle(nums)

	w.Close()
	out, _ := io.ReadAll(r)
	got := string(out)
	want := "2"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHandler_Handle_Medium(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w // Set os.Stdout to the w, the File that we have

	h := Handler{
		Writer:     os.Stdout,
		Calculator: calc.Addition{},
	}
	nums := []string{"10", "15"}
	h.Handle(nums)

	w.Close()
	out, _ := io.ReadAll(r) // read from r, the File connected to w
	got := string(out)
	want := "25"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
