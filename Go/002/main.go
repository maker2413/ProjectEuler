package main

import "fmt"

const (
	to = 4000000
)

func main() {
	var n, m int = 1, 1

	answer := 0

	for m < to {
		m += n
		n = m - n
		if m%2 == 0 {
			answer += m
		}
	}

	fmt.Println(answer)
}
