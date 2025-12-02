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
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		firstHand := strings.Split(scanner.Text(), " ")

		pg := PokerGame{
			p1: ConvertToHand([5]string(firstHand[:5])),
			p2: ConvertToHand([5]string(firstHand[5:])),
		}

		err = pg.CheckWinner()
		if err != nil {
			log.Fatalln(err)
		}

		if pg.Winner == 1 {
			answer++
		}
	}

	fmt.Println("Answer:", answer)
}
