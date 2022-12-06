package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var ct int = 0

func DetermineSignal(a []string) bool {
	if (a[0] != a[1]) && (a[0] != a[2]) && (a[0] != a[3]) && (a[1] != a[2]) && (a[1] != a[3]) && (a[2] != a[3]) {
		return true
	}
	return false
}

var bo bool = false
var p []bool

func DetermineMessage(a []string, i int) {
	if i > 0 {
		for j := i - 1; j >= 0; j-- {
			if a[i] != a[j] {
				//fmt.Println("This happened")
				p = append(p, true)
			} else {
				p = append(p, false)
			}
		}
		DetermineMessage(a, i-1)
	} else {
		bo = true
		for k := range p {
			if !p[k] {
				bo = false
			}
		}
		p = nil
	}
}

func main() {

	f, err := os.Open("signal.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "")
		for i := range a {
			if i+4 < len(a) {
				// if DetermineSignal(a[i:i+4], len(a)) {
				// 	fmt.Println(a[i : i+4])
				// 	break
				// } else {
				// 	ct += 1
				// }
				if !bo {
					DetermineMessage(a[i:i+14], 13)
					ct += 1
				} else {
					fmt.Println(a[i : i+14])
					break
				}

			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ct)
}
