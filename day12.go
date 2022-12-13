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

func pathfinder() {
	start := Point{}
	finish := Point{}
	for a := range grid {
		if grid[a].elev == 0 {
			start = grid[a]
		}
		if grid[a].elev == 27 {
			finish = grid[a]
		}
	}
	fmt.Println(start, " ", finish)
	startString := strconv.Itoa(start.x) + ":" + strconv.Itoa(start.y)
	queue := make(map[string]Point)
	for k, v := range grid {
		queue[k] = v
	}
	visited := map[string]string{}
	fmt.Println(startString)
	visited[startString] = startString
	stepCount := 0
	row := []int{-1, 0, 0, 1}
	col := []int{0, -1, 1, 0}
	failed := false
	current := grid[startString]
	e := ""
	r := ""
	for len(queue) > 0 && !failed {
		stepCount++

		//Check if successful
		if current.elev == finish.elev {
			//End operation
			if current.elev == 27 {
				break
			}
			fmt.Println("End Operation")
			break
		}

		vtrack := len(visited)

		for a := 0; a < 4; a++ {
			e = strconv.Itoa(current.x) + ":" + strconv.Itoa(current.y)
			nextX := current.x + row[a]
			nextY := current.y + col[a]
			//fmt.Println(nextX, nextY)
			r = strconv.Itoa(nextX) + ":" + strconv.Itoa(nextY)

			if isValid(nextX, nextY) && grid[r].elev != 0 { //Path is possible within bounds
				//fmt.Println(grid[r])

				if queue[r].elev == queue[e].elev+1 {
					//fmt.Println(queue[r].elev, queue[e].elev)
					if _, ok := visited[r]; !ok {
						current = queue[r]
						//fmt.Println("Q of R : ", queue[r])
						delete(queue, e)
						visited[e] = e
					}
				}
			}
		}

		if len(visited) == vtrack {
			for a := 0; a < 4; a++ {
				e = strconv.Itoa(current.x) + ":" + strconv.Itoa(current.y)
				nextX := current.x + row[a]
				nextY := current.y + col[a]
				//fmt.Println(nextX, nextY)
				r = strconv.Itoa(nextX) + ":" + strconv.Itoa(nextY)

				if isValid(nextX, nextY) && grid[r].elev != 0 { //Path is possible within bounds
					//fmt.Println(grid[r])

					if queue[r].elev == queue[e].elev {
						//fmt.Println(queue[r].elev, queue[e].elev)
						if _, ok := visited[r]; !ok {
							current = queue[r]
							//fmt.Println("Q of R : ", queue[r])
							delete(queue, e)
							visited[e] = e
						}
					}
				}
			}
		}

		if len(visited) == vtrack {
			for a := 0; a < 4; a++ {
				e = strconv.Itoa(current.x) + ":" + strconv.Itoa(current.y)
				nextX := current.x + row[a]
				nextY := current.y + col[a]
				//fmt.Println(nextX, nextY)
				r = strconv.Itoa(nextX) + ":" + strconv.Itoa(nextY)

				if isValid(nextX, nextY) && grid[r].elev != 0 { //Path is possible within bounds
					//fmt.Println(grid[r])

					if queue[r].elev == queue[e].elev-1 {
						//fmt.Println(queue[r].elev, queue[e].elev)
						if _, ok := visited[r]; !ok {
							current = queue[r]
							//fmt.Println("Q of R : ", queue[r])
							delete(queue, e)
							visited[e] = e
						}
					}
				}
			}
		}

		if len(visited) == vtrack {
			fmt.Println("Dead End")
			break
		}

	}
	fmt.Println(len(visited))
}

func isValid(x, y int) bool {
	return (x >= 0) && (x < maxX) && (y <= 0) && (y > maxY)
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
			if i+1 < maxY {
				if rows[i+1][j] <= rows[i][j]+1 {
					//fmt.Println("Can moved DOWN")
					v += "U"
				}
			}
			if i-1 >= 0 {
				if rows[i-1][j] <= rows[i][j]+1 {
					v += "D"
					//fmt.Println("Can moved UP")
				}
			}
			cartRow = append(cartRow, v)
		}
		cart = append(cart, cartRow)
	}
	for z := range cart {
		fmt.Println(cart[z])
	}
	//Find all Paths
	// Set Step Count to 0
	step := 0
	// Loop through Cart to find available paths
	// If Step Count < Step Count > 0 set Steph Count
	// HOW to REACH the END????
	cx := start.x
	cy := start.y
	failed := false
	for cx != finish.x && cy != finish.y && !failed {
		icart := cart
		if cx == finish.x && cy == finish.y {
			fmt.Println(step)
			break
		}
		n := strings.Split(icart[cy][cx], "")
		for o := range n {
			if n[o] != "." {
				if n[o] == "R" {
					//icart[cy][cx] = strings.Join(removeCart(n, o), " ")
					cx++
					step += 1
					fmt.Println(icart[cy][cx])
					break
				} else if n[o] == "D" {
					//icart[cy][cx] = strings.Join(removeCart(n, o), " ")
					cy++
					step += 1
					fmt.Println(icart[cy][cx])
					break
				} else if n[o] == "L" {
					//icart[cy][cx] = strings.Join(removeCart(n, o), " ")
					cx--
					step += 1
					fmt.Println(icart[cy][cx])
					break
				} else if n[o] == "U" {
					//icart[cy][cx] = strings.Join(removeCart(n, o), " ")
					cy--
					step += 1
					fmt.Println(icart[cy][cx])
					break
				} else {
					step = 0
					failed = true
				}
			} else {
				step = 0
				failed = true
			}
		}
	}
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
		yIndex--

	}
	maxY = yIndex
	maxX = len(rows[0])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pathfinder()

	fmt.Println(maxX, " ", maxY)
}
