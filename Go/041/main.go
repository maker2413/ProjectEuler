package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPandigital(str string) bool {
	digits := make(map[string]int)

	for s := 0; s < len(str); s++ {
		if string(str[s]) == "0" || string(str[s]) == "8" || string(str[s]) == "9" {
			return false
		}

		digits[string(str[s])]++
	}

	if len(digits) == 7 {
		return true
	}

	return false
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Pow(float64(n), .5)); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var answer int

	// since 1+2+3+4+5+6+7+8 = 36 all 8 digit pandigital numbers are divisible by 3
	// similarly 9 digit pandigital numbers add up to 45 which makes them divisible by 3
	for i := 1234567; i <= 7654321; i += 2 {
		if isPandigital(strconv.Itoa(i)) {
			if isPrime(i) && i > answer {
				answer = i
			}
		}
	}

	fmt.Println(answer)
}
