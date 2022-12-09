package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type cor struct {
	x, y float64
}

var H = make(map[float64]float64)
var T = []cor{}
var TT = make(map[string]float64)
var hx float64
var hy float64
var tx float64
var ty float64
var ct int
var dup int

func moveH(d string, i int) {
	for u := 0; u < i; u++ {
		thx := hx
		thy := hy
		if d == "R" {
			hx++
		} else if d == "L" {
			hx--
		} else if d == "U" {
			hy++
		} else if d == "D" {
			hy--
		}
		if checkAdjacent() {

		} else {
			moveT(thx, thy)
		}
		//fmt.Println(thx, " : ", thy)
		fmt.Println("Result : H =(", hx, ",", hy, ") T=(", tx, ",", ty, ")")
	}
}

func checkAdjacent() bool {
	if math.Sqrt(math.Pow(hx-tx, 2)+(math.Pow(hy-ty, 2))) > 1.4142135623730951 {
		return false
	} else {
		//fmt.Println(math.Sqrt(math.Pow(hx-tx, 2) + (math.Pow(hy-ty, 2))))
		return true
	}
}

func moveT(thx float64, thy float64) {
	tx = thx
	ty = thy
	if !checkAdjacent() {
		//fmt.Println("Result : H =(", hx, ",", hy, ") T=(", tx, ",", ty, ")")
		ct += 1
	}
	a := cor{tx, ty}
	T = append(T, a)
}

func checkDups() {
	for i := range T {
		d := -1
		for j := range T {
			if T[i].x == T[j].x && T[i].y == T[j].y {
				d += 1
				s := fmt.Sprintf("%f", T[i].x)
				w := fmt.Sprintf("%f", T[i].y)
				TT[(s + "-" + w)] = 1
				break
			}
		}
		dup += d
	}
}

func main() {

	f, err := os.Open("plank.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		a := strings.Split(scanner.Text(), " ")
		b, _ := strconv.Atoi(a[1])
		fmt.Println("Execute : ", a[0], " ", a[1])
		moveH(a[0], b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Errors in processing : ", ct)
	checkDups()
	fmt.Println("Duplicate locations : ", dup)
	//fmt.Println(TT)
	fmt.Println(len(T) - (dup / 2) + 1)
	fmt.Println(len(TT))
}
