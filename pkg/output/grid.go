package output

import "fmt"

func PrintGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}
