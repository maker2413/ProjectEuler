package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	to = 1000000
)

func main() {
	answer := 1

	var str strings.Builder

	for i := 0; len(str.String()) <= to; i++ {
		str.WriteString(strconv.Itoa(i))
	}

	irrational := str.String()
	for d := 1; d <= to; d *= 10 {
		x, _ := strconv.Atoi(string(irrational[d]))

		answer *= x
	}

	fmt.Println(answer)
}
