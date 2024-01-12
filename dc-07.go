package main

import (
	"fmt"
)

func maxSumNonAdjacent(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	// Special case for lists with only one element
	if length == 1 {
		return nums[0]
	}

	// Initialize an array to store the maximum sum up to each index
	maxSum := make([]int, length)

	// Base cases
	maxSum[0] = nums[0]
	maxSum[1] = max(nums[0], nums[1])

	// Iterate through the rest of the array
	for i := 2; i < length; i++ {
		// Choose the maximum between:
		// 1. The maximum sum at the previous index
		// 2. The sum of the current element and the maximum sum two positions back
		maxSum[i] = max(maxSum[i-1], maxSum[i-2]+nums[i])
	}

	// The final result is the maximum sum at the last index
	return maxSum[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage
	nums := []int{2, 4, 6, 2, 5}
	result := maxSumNonAdjacent(nums)
	fmt.Println("Largest sum of non-adjacent numbers:", result)
}
