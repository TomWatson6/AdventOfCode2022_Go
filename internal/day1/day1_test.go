package day1_test

import (
	"testing"

	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day1"
)

func TestRun(t *testing.T) {
	ans, err := day1.Run("simple_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := 0

	if ans != expected {
		t.Fatalf("incorrect answer, expected: %d, got: %d", expected, ans)
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := day1.Run("simple_input.txt")
		if err != nil {
			b.Fatal("benchmarking invalid on non-working solution for day1")
		}
	}
}
