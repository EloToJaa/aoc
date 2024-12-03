package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func processLine(line string) int {
	sum := 0

	regex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return 0
	}

	matches := regex.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 3 {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				continue
			}

			sum += a * b
		}
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += processLine(line)
		if line == "" {
			break
		}
	}

	fmt.Println(sum)
}
