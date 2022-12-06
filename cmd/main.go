package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day1"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day2"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day3"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day4"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day5"
	"github.com/tomwatson6/AdventOfCode2022_Go/internal/day6"
)

var days = []func(string) string{
	d1, d2, d3, d4, d5, d6,
}

func main() {
	output := make([]string, len(days))
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(days))

	path := "internal/dayX/input.txt"
	for i, d := range days {
		go func(index int, path string, day func(string) string, mu *sync.Mutex, wg *sync.WaitGroup) {
			p := strings.Replace(path, "X", strconv.Itoa(index+1), 1)
			solution := day(p)

			mu.Lock()
			output[index] = solution
			mu.Unlock()

			wg.Done()
		}(i, path, d, mu, wg)
	}

	wg.Wait()

	for _, out := range output {
		fmt.Print(out)
	}
}

func d1(fileName string) string {
	output := "--- Day 1 ---\n"

	d1p1, err := day1.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %d\n", d1p1)

	d1p2, err := day1.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %d\n\n", d1p2)

	return output
}

func d2(fileName string) string {
	output := "--- Day 2 ---\n"

	d2p1, err := day2.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %d\n", d2p1)

	d2p2, err := day2.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %d\n\n", d2p2)

	return output
}

func d3(fileName string) string {
	output := "--- Day 3 ---\n"

	d3p1, err := day3.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %d\n", d3p1)

	d3p2, err := day3.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %d\n\n", d3p2)

	return output
}

func d4(fileName string) string {
	output := "--- Day 4 ---\n"

	d4p1, err := day4.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %d\n", d4p1)

	d4p2, err := day4.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %d\n\n", d4p2)

	return output
}

func d5(fileName string) string {
	output := "--- Day 5 ---\n"

	d5p1, err := day5.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %s\n", d5p1)

	d5p2, err := day5.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %s\n\n", d5p2)

	return output
}

func d6(fileName string) string {
	output := "--- Day 6 ---\n"

	d6p1, err := day6.Part1(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 1: %d\n", d6p1)

	d6p2, err := day6.Part2(fileName)
	if err != nil {
		panic(err)
	}

	output += fmt.Sprintf("Part 2: %d\n\n", d6p2)

	return output
}
