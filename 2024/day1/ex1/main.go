package main

import (
	"fmt"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var input_a, input_b []int

	for {
		var a, b int
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}
		input_a = append(input_a, int(a))
		input_b = append(input_b, int(b))
	}

	sort.Ints(input_a)
	sort.Ints(input_b)

	sum := 0
	for i := 0; i < len(input_a); i++ {
		sum += Abs(input_a[i] - input_b[i])
	}
	fmt.Println(sum)
}
