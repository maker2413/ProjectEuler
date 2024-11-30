package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	factorial := big.NewInt(1)

	for i := 100; i > 0; i-- {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}

	str := factorial.String()

	var answer int

	for i := 0; i < len(str); i++ {
		m, _ := strconv.Atoi(string(str[i]))
		answer += m
	}

	fmt.Println(answer)
}
