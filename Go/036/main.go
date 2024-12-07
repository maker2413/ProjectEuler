package main

import (
	"fmt"
	"strconv"
)

const (
	to = 1000000
)

func palindromeCheck(num string) bool {
	for i := 0; i < len(num)/2; i++ {
		if string(num[i]) != string(num[len(num)-1-i]) {
			return false
		}
	}

	return true
}

func isDoubleBasePalindrome(n int) bool {
	if palindromeCheck(strconv.Itoa(n)) &&
		palindromeCheck(strconv.FormatInt(int64(n), 2)) {
		return true
	}

	return false
}

func main() {
	var answer int

	for i := 1; i < to; i++ {
		if isDoubleBasePalindrome(i) {
			answer += i
		}
	}

	fmt.Println(answer)
}
