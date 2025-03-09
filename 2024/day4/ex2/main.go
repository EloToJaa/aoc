package main

import (
	"bufio"
	"fmt"
	"os"
)

func searchForPattern(input []string, x, y int) int {
	if input[x][y] != 'A' {
		return 0
	}

	a, b, c, d := input[x-1][y-1], input[x-1][y+1], input[x+1][y+1], input[x+1][y-1]

	if a == 'M' && b == 'M' && c == 'S' && d == 'S' {
		return 1
	}
	if a == 'S' && b == 'M' && c == 'M' && d == 'S' {
		return 1
	}
	if a == 'S' && b == 'S' && c == 'M' && d == 'M' {
		return 1
	}
	if a == 'M' && b == 'S' && c == 'S' && d == 'M' {
		return 1
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		input = append(input, line)
	}

	sum := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			sum += searchForPattern(input, i, j)
		}
	}

	fmt.Println(sum)
}
