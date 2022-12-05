package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	id      int
	rations []int
	total   int
}

type Party struct {
	P []Elf
}

var Elves = make(map[int]int)

func NewElf(i int) {
	e := Elf{}
	e.id = i
	ElfParty.P = append(ElfParty.P, e)
}

func AddRation(id int, ct int) {
	fmt.Println(ElfParty.P[id])
	ElfParty.P[id].rations = append(ElfParty.P[id].rations, ct)
}

func SumRation() {
	for j := range ElfParty.P {
		for k := range ElfParty.P[j].rations {
			ElfParty.P[j].total += ElfParty.P[j].rations[k]
		}
	}
}

func FindLargest() {
	b := 0
	i := 0
	for j := range ElfParty.P {
		if ElfParty.P[j].total >= b {
			b = ElfParty.P[j].total
			i = j
		}
	}
	fmt.Println("Elf with the most : ", i, " with this man calories : ", b)
}

func FindTopThree() {
	m := make(map[int]int)
	for j := range ElfParty.P {
		m[j] = ElfParty.P[j].total
	}
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}

var ElfParty Party

func main() {
	f, err := os.Open("elf_rations.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := -1
	for scanner.Scan() {
		if scanner.Text() == "" {
			i++
			NewElf(i)
		} else {
			r, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			AddRation(i, r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	SumRation()
	FindLargest()
	FindTopThree()
}
