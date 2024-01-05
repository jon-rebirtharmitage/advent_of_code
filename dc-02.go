/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	for i := 0; i < len(a); i++ {
		k := 0
		for j := 0; j < len(a); j++ {
			if j != i {
				if k == 0 {
					k = a[j]
				} else {
					k = k * a[j]
				}
			}

		}
		b = append(b, k)
		k = 0
	}
	fmt.Println(b)
}
