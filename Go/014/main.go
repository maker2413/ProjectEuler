package main

import "fmt"

func sequenceLength(n int) int {
	length := 1

	for n != 1 {
		length++
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}

	return length
}

func main() {
	var maxLength, answer int

	for i := 1; i < 1000000; i++ {
		l := sequenceLength(i)
		if l > maxLength {
			maxLength = l
			answer = i
		}
	}

	fmt.Println(answer)
}
