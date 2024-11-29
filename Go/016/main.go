package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	pow = 1000
)

func sumDigits(n *big.Int) int {
	var sum int

	str := n.String()

	for i := 0; i < len(str); i++ {
		s, _ := strconv.Atoi(string(str[i]))
		sum += s
	}

	return sum
}

func main() {
	x := big.NewInt(2)

	x.Exp(x, big.NewInt(pow), big.NewInt(0))

	fmt.Println(sumDigits(x))
}
