package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cycle int
var v int
var signal = make(map[int]int)

func signalCalc(a []int) int {
	res := 0
	for k, v := range signal {
		for i := range a {
			if k == a[i] {
				fmt.Println("For cycle : ", a[i], " Signal Product : ", (k * v))
				res += (k * v)
			}
		}
	}
	return res
}

func main() {

	f, err := os.Open("cycle.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	v = 1
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		if a[0] == "noop" {
			cycle += 1
			signal[cycle] = v
		} else if a[0] == "addx" {
			for c := 0; c < 2; c++ {
				cycle += 1
				signal[cycle] = v
			}
			signal[cycle] = v
			j, _ := strconv.Atoi(a[1])
			v = v + (j)
		}
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(signal)
	fmt.Println("Total Cycle Count: ", cycle)
	s := []int{20, 60, 100, 140, 180, 220}
	fmt.Println("Sum of signals : ", signalCalc(s))
	rowA := [40]string{}
	rowB := [40]string{}
	rowC := [40]string{}
	rowD := [40]string{}
	rowE := [40]string{}
	rowF := [40]string{}
	for k, v := range signal {
		if k <= 40 {
			if v == (k-1) || v+1 == (k-1) || v-1 == (k-1) {
				rowA[k-1] = "#"
			} else {
				rowA[k-1] = "."
			}
		} else if k > 40 && k <= 80 {
			if v == (k-41) || v+1 == (k-41) || v-1 == (k-41) {
				rowB[k-41] = "#"
			} else {
				rowB[k-41] = "."
			}
		} else if k > 80 && k <= 120 {
			if v == (k-81) || v+1 == (k-81) || v-1 == (k-81) {
				rowC[k-81] = "#"
			} else {
				rowC[k-81] = "."
			}
		} else if k > 120 && k <= 160 {
			if v == (k-121) || v+1 == (k-121) || v-1 == (k-121) {
				rowD[k-121] = "#"
			} else {
				rowD[k-121] = "."
			}
		} else if k > 160 && k <= 200 {
			if v == (k-161) || v+1 == (k-161) || v-1 == (k-161) {
				rowE[k-161] = "#"
			} else {
				rowE[k-161] = "."
			}
		} else if k > 200 && k <= 240 {
			if v == (k-201) || v+1 == (k-201) || v-1 == (k-201) {
				rowF[k-201] = "#"
			} else {
				rowF[k-201] = "."
			}
		}
	}
	fmt.Println(rowA)
	fmt.Println(rowB)
	fmt.Println(rowC)
	fmt.Println(rowD)
	fmt.Println(rowE)
	fmt.Println(rowF)
	//fmt.Println(signal)
}
