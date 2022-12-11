package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	id     int
	items  []string
	test   int
	opType string
	op     string
	alpha  int64 //true
	omega  int64 //false
	total  int64
}

var old float64 = 0.0 //tracking last inspection value
var ct int
var monkies []monkey

func monkeyBusiness(z int) {
	for i := 0; i < z; i++ {
		ct += 1
		fmt.Println("Round : ", ct)
		for j := range monkies { // run a round
			h := len(monkies[j].items)
			for k := 0; k < h; k++ {
				monkies[j].total += 1
				//w := big.NewInt(3)
				z := big.NewInt(0)
				var wa string
				e := big.NewInt(0)
				e.SetString(monkies[j].items[0], 10)
				r := big.NewInt(0)
				r.SetString(monkies[j].op, 10)
				q := big.NewInt(0)
				h := strconv.Itoa(h)
				q.SetString(h, 10)
				f := big.NewInt(0)
				if monkies[j].op == "-1" {
					if monkies[j].opType == "mul" {
						z.Mul(e, e)
						//w = monkies[j].items[0] * monkies[j].items[0]
					} else if monkies[j].opType == "add" {
						//w = monkies[j].items[0] + monkies[j].items[0]
						z.Add(e, e)
					}
				} else {
					if monkies[j].opType == "mul" {
						//w = monkies[j].items[0] * monkies[j].op
						z.Mul(e, r)
					} else if monkies[j].opType == "add" {
						//w = monkies[j].items[0] + monkies[j].op
						z.Add(e, r)
					}
				}
				wa = z.String()
				f.DivMod(z, q, f)
				if f.String() != "0" {
					a := remove(monkies[j].items, 0)
					monkies[j].items = a
					monkies[monkies[j].omega].items = append(monkies[monkies[j].omega].items, wa)
				} else {
					a := remove(monkies[j].items, 0)
					monkies[j].items = a
					monkies[monkies[j].alpha].items = append(monkies[monkies[j].alpha].items, wa)
				}
				z = nil
				f = nil
				q = nil
				r = nil
				e = nil
			}
		}
		//fmt.Println("Results of Round ", i+1, " : ", monkies)
	}
}

func checkSumDivison(v string, t int) bool {
	s := strconv.Itoa(t)
	n := big.NewInt(0)
	n.SetString(s, 10)
	m := big.NewInt(0)
	m.SetString(v, 10)
	z := big.NewInt(0)
	z.DivMod(m, n, z)
	//zero := big.NewInt(0)
	//fmt.Println(m, n, z)
	if "0" == z.String() {
		return true
	} else {
		return false
	}
	// if t == 2 {
	// 	a := v[len(v)-1:]
	// 	s, _ := strconv.Atoi(a)
	// 	if s%t == 0 {
	// 		return true
	// 	}
	// } else if t == 7 {
	// 	vt := v
	// 	for len(vt) > 2 {
	// 		h := "-" + vt[len(vt)-1:]
	// 		z := big.NewInt(0)
	// 		q := big.NewInt(2)
	// 		r := big.NewInt(7)
	// 		zero := big.NewInt(0)
	// 		vt = vt[0 : len(vt)-1]
	// 		m := big.NewInt(0)
	// 		m.SetString(vt, 10)
	// 		n := big.NewInt(0)
	// 		n.SetString(h, 10)
	// 		n.Mul(n, q)
	// 		m.Add(m, n)
	// 		z.DivMod(m, r, z)
	// 		if z == zero {
	// 			return true
	// 		}
	// 	}
	// } else if t == 13 {
	// 	p := 1
	// 	s := 0
	// 	n := len(v)
	// 	if n%3 == 1 {
	// 		v = "00" + v
	// 	}

	// 	if n%3 == 2 {
	// 		v = "0" + v
	// 	}
	// 	chunks := Chunks(v, 3)

	// 	for i := len(chunks) - 1; i >= 0; i-- {
	// 		a, _ := strconv.Atoi(chunks[i])
	// 		//b, _ := strconv.Atoi(chunks[i+1])
	// 		s = s + (a * p)
	// 		p = p * -1
	// 	}
	// 	if (s % t) == 0 {
	// 		return true
	// 	}
	// } else if t == 3 {
	// 	n := strings.Split(v, "")
	// 	s := 0
	// 	for i := range n {
	// 		a, _ := strconv.Atoi(n[i])
	// 		s = s + a
	// 	}
	// 	if (s % t) == 0 {
	// 		return true
	// 	}
	// } else if t == 17 {
	// 	vt := v
	// 	for len(vt) > 2 {
	// 		h := vt[len(vt)-1:]
	// 		z := big.NewInt(0)
	// 		q := big.NewInt(5)
	// 		r := big.NewInt(17)
	// 		zero := big.NewInt(0)
	// 		vt = vt[0 : len(vt)-1]
	// 		m := big.NewInt(0)
	// 		m.SetString(vt, 10)
	// 		n := big.NewInt(0)
	// 		n.SetString(h, 10)
	// 		n.Mul(n, q)
	// 		m.Sub(m, n)
	// 		if m == zero {
	// 			return true
	// 		}
	// 		z.DivMod(m, r, z)
	// 		if z == zero {
	// 			return true
	// 		}
	// 	}
	// } else if t == 19 {
	// 	vt := v
	// 	for len(vt) > 2 {
	// 		h := vt[len(vt)-1:]
	// 		z := big.NewInt(0)
	// 		q := big.NewInt(2)
	// 		r := big.NewInt(19)
	// 		zero := big.NewInt(0)
	// 		vt = vt[0 : len(vt)-1]
	// 		m := big.NewInt(0)
	// 		m.SetString(vt, 10)
	// 		n := big.NewInt(0)
	// 		n.SetString(h, 10)
	// 		n.Mul(n, q)
	// 		m.Add(m, n)
	// 		z.DivMod(m, r, z)
	// 		if z == zero {
	// 			return true
	// 		}
	// 	}
	// } else if t == 11 {
	// 	n := strings.Split(v, "")
	// 	odd := 0
	// 	even := 0
	// 	for i := 0; i < len(n); i++ {
	// 		a, _ := strconv.Atoi(n[i])
	// 		if i%2 == 0 {
	// 			even = even + a
	// 		} else {
	// 			odd = odd + a
	// 		}
	// 	}
	// 	if odd >= even {
	// 		if ((odd - even) % t) == 0 {
	// 			return true
	// 		}
	// 	} else {
	// 		if ((even - odd) % t) == 0 {
	// 			return true
	// 		}
	// 	}

	// } else if t == 5 {
	// 	vt := v[len(v)-1:]
	// 	if vt == "0" || vt == "5" {
	// 		return true
	// 	}
	// }
	// return false
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func remove(slice []string, s int64) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {

	f, err := os.Open("monkey.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentMonkey := 0
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if strings.Contains(scanner.Text(), "Monkey") { //new monkey
			a := strings.Split(scanner.Text(), " ")
			b := strings.Split(a[1], "")
			currentMonkey, _ = strconv.Atoi(b[0])
			m := monkey{currentMonkey, nil, 0, "", "", 0, 0, 0}
			monkies = append(monkies, m)
		}
		if strings.Contains(scanner.Text(), "Starting") {
			a := strings.Split(scanner.Text(), ":")
			b := strings.Split(a[1], ",")
			for i := range b {
				//c, err := strconv.ParseInt(strings.TrimSpace(b[i]), 0, 64)
				if err != nil {
					fmt.Println(err)
				}
				monkies[currentMonkey].items = append(monkies[currentMonkey].items, strings.TrimSpace(b[i]))
			}
		}
		if strings.Contains(scanner.Text(), "Operation") {
			if strings.Contains(scanner.Text(), "+") {
				monkies[currentMonkey].opType = "add"
			} else if strings.Contains(scanner.Text(), "*") {
				monkies[currentMonkey].opType = "mul"
			}
			a := strings.Split(scanner.Text(), " ")
			if a[len(a)-1] == "old" {
				monkies[currentMonkey].op = "-1"
			} else {
				//b, _ := strconv.ParseInt(a[len(a)-1], 0, 64)
				monkies[currentMonkey].op = a[len(a)-1]
			}
		}
		if strings.Contains(scanner.Text(), "Test") {
			a := strings.Split(scanner.Text(), " ")
			//b, _ := strconv.ParseInt(a[len(a)-1], 0, 64)
			b, _ := strconv.Atoi(a[len(a)-1])
			monkies[currentMonkey].test = b
		}
		if strings.Contains(scanner.Text(), "true") {
			a := strings.Split(scanner.Text(), " ")
			b, _ := strconv.ParseInt(a[len(a)-1], 0, 64)
			monkies[currentMonkey].alpha = b
		}
		if strings.Contains(scanner.Text(), "false") {
			a := strings.Split(scanner.Text(), " ")
			b, _ := strconv.ParseInt(a[len(a)-1], 0, 64)
			monkies[currentMonkey].omega = b
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(monkies)
	monkeyBusiness(1000)
	for q := range monkies {
		fmt.Println("Total Inspection by monkey ", q, " : ", monkies[q].total)
	}
}
