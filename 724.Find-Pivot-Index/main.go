package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}

	resp := pivotIndex(nums)

	fmt.Println("resp::", resp)
}

func pivotIndex(nums []int) int {
	left, right := 0, 0
	for _, val := range nums {
		right += val
	}

	for ind, val := range nums {
		right -= val
		if left == right {
			return ind
		}
		left += val
	}

	return -1
}
