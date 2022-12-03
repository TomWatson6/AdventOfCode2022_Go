package main

import (
	"fmt"

	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day1"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day2"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day3"
)

func main() {
	d1("internal/day1/input.txt")
	d2("internal/day2/input.txt")
	d3("internal/day3/input.txt")
}

func d1(fileName string) {
	fmt.Println("--- Day 1 ---")

	d1p1, err := day1.Part1(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", d1p1)

	d1p2, err := day1.Part2(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %d\n", d1p2)
	fmt.Println()
}

func d2(fileName string) {
	fmt.Println("--- Day 2 ---")

	d2p1, err := day2.Part1(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", d2p1)

	d2p2, err := day2.Part2(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %d\n", d2p2)
	fmt.Println()
}

func d3(fileName string) {
	fmt.Println("--- Day 3 ---")

	d3p1, err := day3.Part1(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", d3p1)

	d3p2, err := day3.Part2(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %d\n", d3p2)
}
