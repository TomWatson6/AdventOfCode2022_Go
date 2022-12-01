package day1

import (
	"github.com/tomwatson6/AdventOfCode2022_Go/pkg/input"
)

func Run(fileName string) (int, error) {
	_, err := input.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
