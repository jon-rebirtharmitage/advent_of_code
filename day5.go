package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	a := []string{"S", "P", "H", "V", "F", "G"}
	b := []string{"M", "Z", "D", "V", "B", "F", "J", "G"}
	c := []string{"N", "J", "L", "M", "G"}
	d := []string{"P", "W", "D", "V", "Z", "G", "N"}
	e := []string{"B", "C", "R", "V"}
	f := []string{"Z", "L", "W", "P", "M", "S", "R", "V"}
	g := []string{"P", "H", "T"}
	h := []string{"V", "Z", "H", "C", "N", "S", "R", "Q"}
	i := []string{"J", "Q", "V", "P", "G", "L", "F"}

	v, err := os.Open("crane.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer v.Close()

	scanner := bufio.NewScanner(v)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	for scanner.Scan() {
		z := strings.Split(scanner.Text(), " ")
		fmt.Println(z[1], z[3], z[5])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
