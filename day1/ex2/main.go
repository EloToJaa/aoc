package main

import (
	"fmt"
)

func main() {
	var input_a []int
	map_b := make(map[int]int)

	for {
		var a, b int
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}
		input_a = append(input_a, a)
		map_b[b]++
	}

	sum := 0
	for _, a := range input_a {
		sum += a * map_b[a]
	}
	fmt.Println(sum)
}
