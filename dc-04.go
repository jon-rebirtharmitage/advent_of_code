/*
Given an array of integers,
find the first missing positive integer in linear time and constant space. In other words,
find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 4, -1, 1}
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		if a[i]+1 != a[i+1] && a[i]+1 > 0 {
			fmt.Print("missing integer is : ", a[i]+1)
			i = len(a)
		}
	}
}
