package main

import "fmt"

const (
	to = 1000
)

func main() {
	var answer int

	for i := 3; i < to; i++ {
		if i%3 == 0 || i%5 == 0 {
			answer += i
		}
	}

	fmt.Println(answer)
}
