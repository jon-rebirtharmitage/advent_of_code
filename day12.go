package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y    int
	el      string
	elev    int
	path    []int
	parentX int
	parentY int
}

var X []int
var Y []int
var rows [][]int
var maxX int
var maxY int

var grid = make(map[string]Point)

func AssignElevation(s string) int {
	m := []string{"S", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "E"}
	for i := range m {
		if s == m[i] {
			return i
		}
	}
	return -1
}

func isValid(x, y int) bool {
	return (x >= 0) && (x < maxX) && (y >= 0) && (y < maxY)
}

func FindPaths() {
	cart := [][]string{}
	start := Point{}
	finish := Point{}
	for i := range rows {
		cartRow := []string{}
		for j := range rows[i] {
			v := "."
			if rows[i][j] == 0 {
				start.x = j
				start.y = i
			}
			if rows[i][j] == 27 {
				finish.x = j
				finish.y = i
			}
			//check left
			if j-1 >= 0 {
				if rows[i][j-1] <= rows[i][j]+1 {
					v += "L"
					//fmt.Println("Can moved LEFT")
				}
			}
			if j+1 < maxX {
				if rows[i][j+1] <= rows[i][j]+1 {
					v += "R"
					//fmt.Println("Can moved RIGHT")
				}
			}
			if (i + 1) < maxY {
				if rows[i+1][j] <= rows[i][j]+1 {
					//fmt.Println("Can moved DOWN")
					v += "D"
				}
			}
			if i-1 >= 0 {
				if rows[i-1][j] <= rows[i][j]+1 {
					v += "U"
					//fmt.Println("Can moved UP")
				}
			}
			if v == "." {
				if j-1 >= 0 {
					if rows[i][j-1] <= rows[i][j] {
						v += "l"
						//fmt.Println("Can moved LEFT")
					}
				}
				if j+1 < maxX {
					if rows[i][j+1] <= rows[i][j] {
						v += "r"
						//fmt.Println("Can moved RIGHT")
					}
				}
				if (i + 1) < maxY {
					if rows[i+1][j] <= rows[i][j] {
						//fmt.Println("Can moved DOWN")
						v += "d"
					}
				}
				if i-1 >= 0 {
					if rows[i-1][j] <= rows[i][j] {
						v += "u"
						//fmt.Println("Can moved UP")
					}
				}
			}
			cartRow = append(cartRow, v)
		}
		cart = append(cart, cartRow)
	}
	for z := range cart {
		fmt.Println(cart[z])
	}
	//BFS Algorithm
	step := 0
	cx := start.x
	cy := start.y
	visited := map[string]string{}
	queue := []string{}
	queue = append(queue, (strconv.Itoa(cy) + ":" + strconv.Itoa(cx)))
	//visited = append(visited, (strconv.Itoa(cy) + ":" + strconv.Itoa(cx)))
	fmt.Println(finish.x, finish.y)
	for len(queue) > 0 {
		step += 1
		v := queue[0]
		queue = queue[1:]
		cor := strings.Split(v, ":")
		currentX, _ := strconv.Atoi(cor[1])
		currentY, _ := strconv.Atoi(cor[0])
		//fmt.Println("Current Eval : ", cart[currentY][currentX])
		if currentX == finish.x && currentY == finish.y {
			fmt.Println("Success!", currentX, currentY)
			break
		}
		//visited[v] = v
		node := strings.Split(cart[currentY][currentX], "")
		st := len(queue)
		for i := range node {
			if node[i] == "." {
				//do nothing
			}
			//Move Left
			if isValid(currentX, currentY) {
				if node[i] == "D" {
					next := (strconv.Itoa(currentY+1) + ":" + strconv.Itoa(currentX))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "R" {
					next := (strconv.Itoa(currentY) + ":" + strconv.Itoa(currentX+1))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "U" {
					next := (strconv.Itoa(currentY-1) + ":" + strconv.Itoa(currentX))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "L" {
					next := (strconv.Itoa(currentY) + ":" + strconv.Itoa(currentX-1))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				}
			} else {

			}
		}
		if len(queue) == st && isValid(currentX, currentY) {
			//fmt.Println("No Good Path")
			for i := range node {
				if node[i] == "l" {
					next := (strconv.Itoa(currentY) + ":" + strconv.Itoa(currentX-1))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "r" {
					next := (strconv.Itoa(currentY) + ":" + strconv.Itoa(currentX+1))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "u" {
					next := (strconv.Itoa(currentY-1) + ":" + strconv.Itoa(currentX))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				} else if node[i] == "d" {
					next := (strconv.Itoa(currentY+1) + ":" + strconv.Itoa(currentX))
					if _, ok := visited[v]; !ok {
						queue = append([]string{next}, queue...)
					}
				}
			}
		}
		visited[v] = v
	}
	fmt.Println(len(visited) - 1)
}

func removeCart(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {

	f, err := os.Open("path.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	yIndex := 0
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		s := strings.Split(scanner.Text(), "")
		cRow := []int{}
		for i := range s {
			p := Point{i, yIndex, s[i], AssignElevation(s[i]), nil, 0, 0}
			s := strconv.Itoa(i) + ":" + strconv.Itoa(yIndex)
			grid[s] = p
			cRow = append(cRow, p.elev)
			X = append(X, i)
			Y = append(Y, yIndex)

		}
		rows = append(rows, cRow)
		yIndex++

	}
	maxY = yIndex
	maxX = len(rows[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//pathfinder()
	FindPaths()

	fmt.Println(maxX, " ", maxY)
}
