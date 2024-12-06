package main

import (
	"fmt"
	"strconv"
)

// found at: https://github.com/TheAlgorithms/Go/blob/82b6e969824f0ee9c36c875857d44563bdfbcfc7/math/gcd/gcd.go
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func digitCancelling(n, m int) bool {
	strN := strconv.Itoa(n)
	strM := strconv.Itoa(m)

	n0, _ := strconv.Atoi(string(strN[0]))
	n1, _ := strconv.Atoi(string(strN[1]))
	m0, _ := strconv.Atoi(string(strM[0]))
	m1, _ := strconv.Atoi(string(strM[1]))

	if n0 == m0 {
		if float64(n1)/float64(m1) == float64(n)/float64(m) {
			return true
		}
	} else if n0 == m1 {
		if float64(n1)/float64(m0) == float64(n)/float64(m) {
			return true
		}
	} else if n1 == m0 {
		if float64(n0)/float64(m1) == float64(n)/float64(m) {
			return true
		}
	} else if n1 == m1 && n1 != 0 && m1 != 0 {
		if float64(n0)/float64(m0) == float64(n)/float64(m) {
			return true
		}
	}

	return false
}

func main() {
	var answer, nProd, dProd int = 0, 1, 1

	for a := 11; a < 99; a++ {
		for b := a + 1; b <= 99; b++ {
			if digitCancelling(a, b) {
				nProd *= a
				dProd *= b
			}
		}
	}

	// Product of denominators divided by gcd of product of all fractions
	answer = int(int64(dProd) / gcd(int64(nProd), int64(dProd)))

	fmt.Println(answer)
}
