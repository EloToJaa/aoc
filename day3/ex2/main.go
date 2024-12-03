package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func lookForMul(input string) int {
	if len(input) < 8 {
		return 0
	}

	start := "mul("
	for i := 0; i < len(start); i++ {
		if input[i] != start[i] {
			return 0
		}
	}

	till := strings.Index(input[:len(start)+4], ",")
	if till == -1 {
		return 0
	}

	digits := "0123456789"
	for i := len(start); i < till; i++ {
		if !strings.Contains(digits, string(input[i])) {
			return 0
		}
	}

	num1, err := strconv.Atoi(input[len(start):till])
	if err != nil {
		return 0
	}

	newStart := till + 1

	till2 := strings.Index(input[newStart:], ")")
	if till2 == -1 {
		return 0
	}

	for i := newStart; i < newStart+till2; i++ {
		if !strings.Contains(digits, string(input[i])) {
			return 0
		}
	}

	num2, err := strconv.Atoi(input[newStart : newStart+till2])
	if err != nil {
		return 0
	}

	return num1 * num2
}

func lookFor(input string, substr string) bool {
	if len(input) < len(substr) {
		return false
	}
	for i := 0; i < len(substr); i++ {
		if input[i] != substr[i] {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	input := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		input += line + "\n"
	}

	enabled := true
	for i := 0; i < len(input)-4; i++ {
		maxLen := int(math.Min(float64(i+12), float64(len(input)-1)))

		if enabled && lookFor(input[i:maxLen], "don't()") {
			enabled = false
			i += 6
			continue
		}
		if !enabled && lookFor(input[i:maxLen], "do()") {
			enabled = true
			i += 3
			continue
		}

		num := lookForMul(input[i:maxLen])
		if num > 0 {
			i += 7
		}

		if enabled {
			sum += num
		}
	}

	fmt.Println(sum)
}
