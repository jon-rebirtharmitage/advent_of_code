package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	alpha   []string
	beta    []string
	gamma   [][]int
	delta   [][]int
	compare bool
}

var packets []packet
var input []string

func breakoutComponents(p packet, i int) {
	currentSliceA := []int{}
	currentSliceB := []int{}
	for a := range p.alpha {
		if p.alpha[a] == "[" && len(currentSliceA) != 0 {
			p.gamma = append(p.gamma, currentSliceA)
			currentSliceA = []int{}
		} else if p.alpha[a] == " " {
			//do nothing
		} else if p.alpha[a] == "]" {
			p.gamma = append(p.gamma, currentSliceA)
			currentSliceA = []int{}
		} else {
			h, err := strconv.Atoi(p.alpha[a])
			if err != nil {

				continue
			} else {
				currentSliceA = append(currentSliceA, h)
			}
		}
	}
	for b := range p.beta {
		if p.beta[b] == "[" && len(currentSliceB) != 0 {
			p.delta = append(p.delta, currentSliceB)
			currentSliceB = []int{}
		} else if p.beta[b] == " " {
			//do nothing
		} else if p.beta[b] == "]" {
			p.delta = append(p.delta, currentSliceB)
			currentSliceB = []int{}
		} else {
			h, err := strconv.Atoi(p.beta[b])
			if err != nil {
				continue
			}
			currentSliceB = append(currentSliceB, h)
		}
	}
	packets[i] = p
}

func compareComponents(p packet, i int) {
	for c := range p.delta {
		if len(p.gamma[c]) < len(p.delta[c]) {
			p.compare = true
			break
		}
		if len(p.gamma[c]) > len(p.delta[c]) {
			p.compare = false
			break
		}
	}
	for a := range p.delta {
		if a < len(p.gamma) {
			for b := range p.delta[a] {
				if b < len(p.gamma[a]) {
					if p.gamma[a][b] < p.delta[a][b] {
						p.compare = true
						break
					} else if p.gamma[a][b] > p.delta[a][b] {
						p.compare = false
						break
					}
				}
			}
		}
	}

	packets[i] = p
}

func main() {

	f, err := os.Open("packet.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			//newline - ignore
		} else {
			input = append(input, scanner.Text())
		}
	}

	for i := 0; i < len(input); i += 2 {
		p := packet{strings.Split(input[i], ""), strings.Split(input[i+1], ""), nil, nil, false}
		packets = append(packets, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for a := range packets {
		breakoutComponents(packets[a], a)
	}

	for c := range packets {
		compareComponents(packets[c], c)
	}
	ct := 0
	for b := range packets {
		fmt.Println("++++++++++++++")
		fmt.Println(packets[b].gamma)
		fmt.Println(packets[b].delta)
		fmt.Println(packets[b].compare)
		if packets[b].compare {
			ct += b + 1
		}
		fmt.Println("++++++++++++++")
	}
	fmt.Println("Kiss my ass motherfucker : ", ct)
}
