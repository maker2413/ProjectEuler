package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var answer int

	file, err := os.Open("poker.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		firstHand := strings.Split(scanner.Text(), " ")

		pg := PokerGame{
			p1: ConvertToHand([5]string(firstHand[:5])),
			p2: ConvertToHand([5]string(firstHand[5:])),
		}

		pg.CheckWinner()

		if pg.Winner == 1 {
			answer++
		}
	}

	fmt.Println(answer)
}
