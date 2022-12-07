package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	id     string
	parent string   //Folder that contains this folder
	dirs   []string //All directories in this directory id being name
	cont   []int    //Files in the directory
	total  int      //Total size of this directory
}

func main() {
	m := map[string]Dir{}
	current := "/"
	parent := "ROOT"
	ct := 0
	driveSum := 0
	//foo := false
	f, err := os.Open("filespace.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ct += 1
		if parent == "" {
			//parent = "/"
			fmt.Println("ERROR ")
			break
		}
		if strings.Contains(scanner.Text(), "$") {
			if strings.Contains(scanner.Text(), "$") {
				if scanner.Text() == "$ cd .." {
					a := strings.Split(current, "/")
					current = ""
					parent = ""
					u := false
					for i := 0; i < len(a)-2; i++ {
						if a[i] != "" {
							if !u {
								current += "/" + a[i] + "/"
								u = true
							} else {
								current += a[i] + "/"
							}
						}
					}
					b := strings.Split(current, "/")
					k := false
					for j := 0; j < len(b)-2; j++ {
						if a[j] != "" {
							if !k {
								parent += "/" + b[j] + "/"
								k = true
							} else {
								parent += b[j] + "/"
							}
						}
					}
					if parent == "" {
						parent = "/"
					}
					if current == "" {
						current = "/"
					}
				} else if strings.Contains(scanner.Text(), "$ ls") {
					//do nothing
				} else if strings.Contains(scanner.Text(), "$ cd ") && scanner.Text() != "$ cd .." {
					a := strings.Split(scanner.Text(), " ")
					parent = current
					if a[2] == "/" {
						current = a[2]
					} else {
						current = parent + a[2] + "/"
					}

				}
			}
		} else if strings.Contains(scanner.Text(), "dir ") {
			a := strings.Split(scanner.Text(), " ")
			m[current+a[1]+"/"] = Dir{current + a[1], parent, nil, nil, 0}
		} else if !strings.Contains(scanner.Text(), "$") && !strings.Contains(scanner.Text(), "dir ") {
			a := strings.Split(scanner.Text(), " ")
			b, _ := strconv.Atoi(a[0])
			driveSum += b
			d := Dir{m[current].id, m[current].parent, m[current].dirs, m[current].cont, (m[current].total + b)}
			m[current] = d
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	final := 0
	for k, _ := range m {
		for b, _ := range m {
			if strings.Contains(b, k) && b != k {
				d := Dir{m[k].id, m[k].parent, m[k].dirs, m[k].cont, (m[k].total + m[b].total)}
				m[k] = d
			}
		}
	}

	for k, _ := range m {
		if m[k].total <= 100000 {
			//fmt.Println("Found one", m[k])
			final += m[k].total
		}
	}

	for q, _ := range m {
		if m[q].total >= 19731435 {
			fmt.Println("Found one", m[q])
			//final += m[k].total
		}
	}

	fmt.Println(final)
	//fmt.Println(ct)
	fmt.Println(driveSum)
	//fmt.Println(len(m))
}
