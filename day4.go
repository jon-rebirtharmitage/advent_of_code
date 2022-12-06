package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ct int = 0

func ExpandGroup(a string) {
	b := strings.Split(a, ",")
	w := strings.Split(b[0], "-")
	x := strings.Split(b[1], "-")
	m, _ := strconv.Atoi(w[0])
	n, _ := strconv.Atoi(w[1])
	o, _ := strconv.Atoi(x[0])
	p, _ := strconv.Atoi(x[1])
	if (m >= o && n <= p) || (o >= m && p <= n) {
		//fmt.Println("Success")
		//fmt.Println(w, x)
		//ct += 1
	}
	if (o <= n && o >= m) || (m <= p && m >= o) {
		fmt.Println(w, x)
		ct += 1
	}
}

func main() {

	f, err := os.Open("groups.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ExpandGroup(scanner.Text())
	}

	fmt.Println(ct)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
