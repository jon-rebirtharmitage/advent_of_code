package main

import (
	"fmt"
)

func main() {
	a := []int{10, 15, 3, 7}
	k := 17
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i]+a[j] == k {
				fmt.Println("True")
				i = len(a)
				j = len(a)
			}
		}
	}
}
