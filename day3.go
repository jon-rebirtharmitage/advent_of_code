package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var total int = 0
var subtotal int = 0

func AssignValue(a string) int {
	res := 0
	m := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := range m {
		if a == m[i] {
			res = i + 1
		}
	}
	fmt.Println(res)
	return res
}

func FindShared(a []string, b []string) {
	z := 0
	for i := range a {
		if contains(b, a[i]) && z == 0 {
			z += 1
			total += AssignValue(a[i])
		}
	}
	fmt.Println(total)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GroupUp(a []string) {
	x := strings.Split(a[0], "")
	y := strings.Split(a[1], "")
	z := strings.Split(a[2], "")
	p := 0
	for i := range x {
		if contains(y, x[i]) && contains(z, x[i]) && p == 0 {
			fmt.Println(x[i])
			p += 1
			total += AssignValue(x[i])
		}
	}
}

func main() {

	f, err := os.Open("rucksacks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	com := []string{}
	for scanner.Scan() {
		l := 0
		a := scanner.Text()
		b := strings.Split(a, "")
		c := []string{}
		d := []string{}
		if len(b)%2 == 0 {
			l = len(b) / 2
		}
		for i := 0; i < l; i++ {
			c = append(c, b[i])
		}
		for j := l; j < len(b); j++ {
			d = append(d, b[j])
		}
		com = append(com, a)

		if len(com) >= 3 {
			GroupUp(com)
			com = nil
		}
		//FindShared(c, d)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
