package main

import "testing"

func TestSolve(t *testing.T) {
	example := "98+76-5+43-2-10"
	want := 200
	got, err := Solve(example)
	if got != want || err != nil {
		t.Fatalf("Got %v, want %v\n", got, want)
	}
}
