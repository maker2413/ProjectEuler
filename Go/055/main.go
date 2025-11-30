package main

import (
	"fmt"
	"strconv"
)

const (
	limit      = 10000
	iterations = 49
)

func palindromeCheck(num string) bool {
	for i := 0; i < len(num)/2; i++ {
		if string(num[i]) != string(num[len(num)-1-i]) {
			return false
		}
	}

	return true
}

func makePalindrome(num int) int {
	s := strconv.Itoa(num)
	var palin string

	for i := range s {
		palin += string(s[len(s)-i-1])
	}

	num, _ = strconv.Atoi(palin)

	return num
}

func lychrelCheck(num int) bool {
	lychrel := num

	for range iterations {
		lychrel += makePalindrome(lychrel)

		if palindromeCheck(strconv.Itoa(lychrel)) {
			break
		}
	}

	if !palindromeCheck(strconv.Itoa(lychrel)) {
		return true
	}

	return false
}

func main() {
	var answer int

	for i := 11; i < limit; i++ {
		if lychrelCheck(i) {
			answer++
		}
	}

	fmt.Println(answer)
}
