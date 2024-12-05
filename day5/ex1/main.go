package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkNums(nums []int, order [][]int) int {
	for i, num := range nums {
		for j := 1; j < len(order); j++ {
			a, b := order[j][0], order[j][1]
			if num == a {
				for k := 0; k < i; k++ {
					if nums[k] == b {
						return 0
					}
				}
			}
		}
	}
	return nums[len(nums)/2]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	order := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		data := make([]int, 2)
		fmt.Sscanf(line, "%d|%d", &data[0], &data[1])
		order = append(order, []int{data[0], data[1]})
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		printOrder := strings.Split(line, ",")
		nums := make([]int, 0)
		for _, order := range printOrder {
			orderNum, err := strconv.Atoi(order)
			if err != nil {
				continue
			}
			nums = append(nums, orderNum)
		}
		sum += checkNums(nums, order)
	}

	fmt.Println(sum)
}
