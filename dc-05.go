package main

import (
	"fmt"
)

type cons struct {
	a, b int
}

func (a *cons) car() int {
	return a.a
}

func (a *cons) cdr() int {
	return a.b
}

func main() {
	alpha := cons{1, 3}
	fmt.Println(alpha)
	fmt.Println(alpha.car())
	fmt.Println(alpha.cdr())
}
