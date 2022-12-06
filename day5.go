package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvertIt(a string) string {
	switch a {
	case "1":
		return "a"
	case "2":
		return "b"
	case "3":
		return "c"
	case "4":
		return "d"
	case "5":
		return "e"
	case "6":
		return "f"
	case "7":
		return "g"
	case "8":
		return "h"
	case "9":
		return "i"
	}
	return ""
}

func main() {

	crane := [][]string{}
	// a := []string{"S", "P", "H", "V", "F", "G"}
	// b := []string{"M", "Z", "D", "V", "B", "F", "J", "G"}
	// c := []string{"N", "J", "L", "M", "G"}
	// d := []string{"P", "W", "D", "V", "Z", "G", "N"}
	// e := []string{"B", "C", "R", "V"}
	// f := []string{"Z", "L", "W", "P", "M", "S", "R", "V"}
	// g := []string{"P", "H", "T"}
	// h := []string{"V", "Z", "H", "C", "N", "S", "R", "Q"}
	// i := []string{"J", "Q", "V", "P", "G", "L", "F"}
	a := []string{"G", "F", "V", "H", "P", "S"}
	b := []string{"G", "J", "F", "B", "V", "D", "Z", "M"}
	c := []string{"G", "M", "L", "J", "N"}
	d := []string{"N", "G", "Z", "V", "D", "W", "P"}
	e := []string{"V", "R", "C", "B"}
	f := []string{"V", "R", "S", "M", "P", "W", "L", "Z"}
	g := []string{"T", "H", "P"}
	h := []string{"Q", "R", "S", "N", "C", "H", "Z", "V"}
	i := []string{"F", "L", "G", "P", "V", "Q", "J"}
	crane = append(crane, a)
	crane = append(crane, b)
	crane = append(crane, c)
	crane = append(crane, d)
	crane = append(crane, e)
	crane = append(crane, f)
	crane = append(crane, g)
	crane = append(crane, h)
	crane = append(crane, i)

	v, err := os.Open("crane.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer v.Close()

	scanner := bufio.NewScanner(v)

	for scanner.Scan() {
		z := strings.Split(scanner.Text(), " ")
		//Get the last one of Z[3] and move it to Z[5], loop z[1] times
		move, _ := strconv.Atoi(z[1])
		from, _ := strconv.Atoi(z[3])
		dest, _ := strconv.Atoi(z[5])
		me := []string{}
		me = crane[from-1][len(crane[from-1])-move : len(crane[from-1])]
		fmt.Println(me)
		fmt.Println("Before : ", crane[from-1], " ", crane[dest-1])
		crane[from-1] = crane[from-1][0 : len(crane[from-1])-(move)]
		// for q := 0; q < move; q++ {
		// 	//fmt.Println("Before : ", crane[from-1], " ", crane[dest-1])
		// 	//crane[dest-1] = append(crane[dest-1], crane[from-1][len(crane[from-1])-1])
		// 	crane[from-1] = crane[from-1][:len(crane[from-1])-1]
		// 	//fmt.Println("After : ", crane[from-1], " ", crane[dest-1])
		// }
		crane[dest-1] = append(crane[dest-1], me...)
		fmt.Println("After : ", crane[from-1], " ", crane[dest-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(crane[0], crane[1], crane[2], crane[3], crane[4], crane[5], crane[6], crane[7], crane[8])
}
