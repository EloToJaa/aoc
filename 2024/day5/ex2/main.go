package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkNums(nums []int, order [][]int) bool {
	for i, num := range nums {
		for j := 1; j < len(order); j++ {
			a, b := order[j][0], order[j][1]
			if num == a {
				for k := 0; k < i; k++ {
					if nums[k] == b {
						return false
					}
				}
			}
		}
	}
	return true
}

func correctOrder(nums []int, order [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		for _, o := range order {
			a, b := o[0], o[1]
			if nums[i] == a && nums[j] == b {
				return false
			}
			if nums[i] == b && nums[j] == a {
				return true
			}
		}
		return nums[i] < nums[j]
	})

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
		correct := checkNums(nums, order)
		if correct {
			continue
		}
		sum += correctOrder(nums, order)
	}

	fmt.Println(sum)
}
