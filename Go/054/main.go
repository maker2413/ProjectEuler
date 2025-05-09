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
	scanner.Scan()

	firstHand := strings.Split(scanner.Text(), " ")

	fmt.Println(firstHand)

	P1 := ConvertToHand([5]string(firstHand[:5]))
	P2 := ConvertToHand([5]string(firstHand[5:]))

	P1.Print()
	P2.Print()

	if P1.IsRoyalFlush() {
		answer++
	}

	fmt.Println(answer)
}
