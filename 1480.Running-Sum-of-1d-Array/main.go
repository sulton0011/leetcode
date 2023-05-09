package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	resp := runningSum(nums)

	fmt.Println("resp::", resp)
}

func runningSum(nums []int) []int {
	var end int
	for ind, val := range nums {
		end += val
		nums[ind] = end
	}

	return nums
}
