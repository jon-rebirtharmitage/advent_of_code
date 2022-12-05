package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
A = Rock
B = Paper
C = Sissors
X = Rock
B = Paper
C = Sissors
Rock > Sissors
Paper > Rock
Sissors > Paper
*/
type Record struct {
	Player   string
	Opp      string
	Score    int
	OppScore int
	winner   int
}

var R []Record
var Rounds []int
var OppRounds []int
var TotalScore int
var OppTotalScore int
var Round int

func DetermineScore(a string, b string) {
	score := 0
	oppScore := 0
	winner := 0
	if b == "X" { //Lose
		if a == "A" {
			b = "Z"
		} else if a == "B" {
			b = "X"
		} else if a == "C" {
			b = "Y"
		}
	} else if b == "Y" { //Tie
		if a == "A" {
			b = "X"
		} else if a == "B" {
			b = "Y"
		} else if a == "C" {
			b = "Z"
		}
	} else if b == "Z" { //Win
		if a == "A" {
			b = "Y"
		} else if a == "B" {
			b = "Z"
		} else if a == "C" {
			b = "X"
		}
	}

	if a == "A" { //Rock
		oppScore += 1 //Selected Rock
		if b == "X" { //Tie
			score += 1 //Selected Rock
			winner = 2 //Draw
			score += 3
			oppScore += 3
		} else if b == "Y" { //Win
			score += 2 //Selected Paper
			winner = 1 //Player
			score += 6
		} else if b == "Z" { //Loss
			score += 3 //Selected Sissor
			winner = 0 //Opp
			oppScore += 6
		}
	} else if a == "B" { //Selected Paper
		oppScore += 2 //Selected Paper
		if b == "X" { //Loss
			score += 1 //Selected Rock
			winner = 0 //Opp
			oppScore += 6
		} else if b == "Y" { //Selcted Paper
			score += 2 //Selected Paper
			winner = 2 //Tie
			score += 3
			oppScore += 3
		} else if b == "Z" { //Selected Sissor
			score += 3 //Selected Sissor
			winner = 1 //Player
			score += 6
		}
	} else if a == "C" {
		oppScore += 3 //Selected Sissor
		if b == "X" { //Win
			score += 1 //Selected Rock
			winner = 1 //Player
			score += 6
		} else if b == "Y" { //Selcted Paper
			score += 2 //Selected Paper
			winner = 0 //Loss
			oppScore += 6
		} else if b == "Z" { //Selected Sissor
			score += 3 //Selected Sissor
			winner = 2 //tie
			score += 3
			oppScore += 3
		}
	}
	Rounds = append(Rounds, score)
	OppRounds = append(OppRounds, oppScore)
	TotalScore += score
	OppTotalScore += oppScore
	R = append(R, Record{a, b, score, oppScore, winner})
}

func main() {

	f, err := os.Open("strategy.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		Re := Record{}
		a := scanner.Text()
		b := strings.Split(a, " ")
		Re.Player = b[0]
		Re.Opp = b[1]
		DetermineScore(Re.Player, Re.Opp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(TotalScore)
}
