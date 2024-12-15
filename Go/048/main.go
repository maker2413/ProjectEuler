package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	digits = 10
)

func lastTenDigits(s, value string) uint64 {
	var i uint64

	if len(s) >= digits {
		s = string(s[len(s)-digits:])

		if string(s[0]) == "0" {
			s = "1" + s
		}

		i, _ = strconv.ParseUint(s, 10, 64)
	} else {
		i, _ = strconv.ParseUint(s, 10, 64)
	}

	return i
}

func main() {
	var answer uint64

	x := big.NewInt(1)

	for i := 1000; i > 0; i-- {
		x = big.NewInt(int64(i))
		x.Exp(x, big.NewInt(int64(i)), nil)

		answer += lastTenDigits(x.String(), "x")
		answer = lastTenDigits(fmt.Sprint(answer), "answer")
	}

	fmt.Println(answer)
}
