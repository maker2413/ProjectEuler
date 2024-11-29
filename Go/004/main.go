package main

import (
	"fmt"
	"strconv"
)

func palindromeCheck(n int) bool {
	num := strconv.Itoa(n)

	for i := 0; i < len(num)/2; i++ {
		if string(num[i]) != string(num[len(num)-1-i]) {
			return false
		}
	}

	return true
}

func main() {
	var answer int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			n := i * j
			if palindromeCheck(n) && n > answer {
				answer = n
			}
		}
	}

	fmt.Println(answer)
}
