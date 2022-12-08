package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findVisible(input [][]int) (int, int) {
	s := 0
	S := 0
	//Get perimeter resolved
	for i := range input {
		for j := range input[i] {
			//fmt.Println("Coordinate : ", i, ",", j, "  Value : ", input[i][j])
			if i == 0 || i == (len(input[i])-1) || j == 0 || j == (len(input[i])-1) { //
				s++
			} else {
				x := i
				y := j
				t := true
				U := []bool{}
				D := []bool{}
				L := []bool{}
				R := []bool{}
				//check up
				for t {
					for u := (x - 1); u >= 0; u-- {
						//fmt.Println("Compare : ", input[x][y], " to ", input[u][y])
						if input[x][y] > input[u][y] {
							U = append(U, true)
						}
						if input[x][y] <= input[u][y] {
							t = false
							U = append(U, false)
							break
						}
					}
					for d := (x + 1); d <= len(input[i])-1; d++ {
						//fmt.Println("Compare : ", input[x][y], " to ", input[d][y])
						if input[x][y] > input[d][y] {
							D = append(D, true)
						}
						if input[x][y] <= input[d][y] {
							t = false
							D = append(D, false)
							break
						}
					}
					for l := (y - 1); l >= 0; l-- {
						//fmt.Println("Compare : ", input[x][y], " to ", input[x][l])
						if input[x][y] > input[x][l] {
							t = false
							L = append(L, true)
						}
						if input[x][y] <= input[x][l] {
							t = false
							L = append(L, false)
							break
						}
					}
					for r := (y + 1); r <= (len(input) - 1); r++ {
						//fmt.Println("Compare : ", input[x][y], " to ", input[x][r])
						if input[x][y] > input[x][r] {
							t = false
							R = append(R, true)
						}
						if input[x][y] <= input[x][r] {
							t = false
							R = append(R, false)
							break
						}
					}
				}
				if Check(U) && Check(D) && Check(L) && Check(R) {

				} else {
					s += 1

				}
				scenic := len(U) * len(D) * len(L) * len(R)
				if scenic > S {
					S = scenic
				}

			}
		}
	}
	return s, S
}

func Check(s []bool) bool {
	for _, v := range s {
		if !v {
			return true
		}
	}

	return false
}

func main() {

	f, err := os.Open("eight.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	rows := [][]int{}
	iter := 0
	for scanner.Scan() {
		row := []int{}
		//fmt.Println(scanner.Text())
		a := strings.Split(scanner.Text(), "")
		for i := range a {
			b, _ := strconv.Atoi(a[i])
			row = append(row, b)
		}
		rows = append(rows, row)
		iter += 1
	}
	//fmt.Println(rows)
	result, scenic := findVisible(rows)
	fmt.Println(result, scenic)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
