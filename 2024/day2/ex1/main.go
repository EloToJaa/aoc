package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkInput(input []int) bool {
	ascending := true
	for i := 1; i < len(input); i++ {
		if input[i]-input[i-1] > 3 || input[i]-input[i-1] < 1 {
			ascending = false
		}
	}

	descending := true
	for i := 1; i < len(input); i++ {
		if input[i-1]-input[i] > 3 || input[i-1]-input[i] < 1 {
			descending = false
		}
	}

	return ascending || descending
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		strNumbers := strings.Fields(line)
		var input []int

		for _, strNumber := range strNumbers {
			num, err := strconv.Atoi(strNumber)
			if err != nil {
				break
			}
			input = append(input, num)
		}

		if checkInput(input) {
			sum++
		}

		if len(input) == 0 {
			break
		}

		input = make([]int, 0)
	}

	fmt.Println(sum)
}
