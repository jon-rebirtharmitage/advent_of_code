package main

import (
	"fmt"
)

type cons struct {
	a, b int
}

func (a *cons) con() {
	fmt.Println("Cons Pair : ", a.a, ":", a.b)
}

func (a *cons) car() int {
	return a.a
}

func (a *cons) cdr() int {
	return a.b
}

func main() {
	alpha := cons{1, 3}
	alpha.con()
	fmt.Println(alpha.car())
	fmt.Println(alpha.cdr())
}
