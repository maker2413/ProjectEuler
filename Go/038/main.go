package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	to = 1000000
)

func isPandigital(str string) bool {
	digits := make(map[string]int)

	for s := 0; s < len(str); s++ {
		if string(str[s]) == "0" {
			return false
		}

		digits[string(str[s])]++
	}

	if len(digits) == 9 {
		return true
	}

	return false
}

func main() {
	var answer int

	for i := 1; i < to; i++ {
		var str strings.Builder

		for j := 1; len(str.String()) < 9; j++ {
			str.WriteString(strconv.Itoa(i * j))

			if len(str.String()) == 9 {
				if isPandigital(str.String()) {
					x, _ := strconv.Atoi(str.String())

					if x > answer {
						answer = x
					}
				}
			}
		}
	}

	fmt.Println(answer)
}
