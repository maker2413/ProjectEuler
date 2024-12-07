package main

import "fmt"

func pentagonalNumber(n int) int {
	return (n * (3*n - 1)) / 2
}

func main() {
	var answer int
	var pentaNums []int

	checkPentaNums := make(map[int]int)

	for i := 1; i <= 10000; i++ {
		pentaNums = append(pentaNums, pentagonalNumber(i))
		checkPentaNums[pentagonalNumber(i)]++
	}

	answer = pentaNums[len(pentaNums)-1]

	for a := 0; a < len(pentaNums)-1; a++ {
		for b := a + 1; b < len(pentaNums); b++ {
			if checkPentaNums[pentaNums[a]+pentaNums[b]] == 1 &&
				checkPentaNums[pentaNums[b]-pentaNums[a]] == 1 {
				n := pentaNums[b] - pentaNums[a]

				if n < answer {
					answer = n
				}
			}
		}
	}

	fmt.Println(answer)
}
