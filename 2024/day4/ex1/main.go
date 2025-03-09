package main

import (
	"bufio"
	"fmt"
	"os"
)

func searchRecursive(input []string, lookFor string, x, y, i, j, k int) int {
	if x < 0 || y < 0 || x >= len(input) || y >= len(input[x]) {
		return 0
	}

	if input[x][y] != lookFor[k] {
		return 0
	}

	if k == len(lookFor)-1 {
		return 1
	}

	return searchRecursive(input, lookFor, x+i, y+j, i, j, k+1)
}

func startRecursive(input []string, x, y int) int {
	lookFor := "XMAS"

	if input[x][y] != lookFor[0] {
		return 0
	}

	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			sum += searchRecursive(input, lookFor, x+i, y+j, i, j, 1)
		}
	}

	return sum
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
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			sum += startRecursive(input, i, j)
		}
	}

	fmt.Println(sum)
}
