package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	to = 10000000
)

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

func genRegex(str string) []string {
	var regexPatterns []string

	for d := 0; d <= 9; d++ {
		if strings.Contains(str, strconv.Itoa(d)) {
			regexPatterns = append(regexPatterns,
				strings.ReplaceAll(str, strconv.Itoa(d), "."))
		}
	}

	return regexPatterns
}

func substituteDigits(str string, primeMap map[int][]string) []int {
	var primes []int

	for i := 0; i <= 9; i++ {
		s := strings.ReplaceAll(str, ".", strconv.Itoa(i))
		num, _ := strconv.Atoi(string(s))

		if len(strconv.Itoa(num)) == len(str) && len(primeMap[num]) > 0 {
			primes = append(primes, num)
		}
	}

	return primes
}

func main() {
	var answer int
	var primes []int
	primeMap := make(map[int][]string)

	for i := 11; i < to; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
			primeMap[i] = genRegex(strconv.Itoa(i))
		}
	}

substitutions:
	// https://www.ardanlabs.com/blog/2013/11/label-breaks-in-go.html
	for prime := range primes {
		regex := genRegex(strconv.Itoa(prime))
		for s := 0; s < len(regex); s++ {
			variations := substituteDigits(regex[s], primeMap)
			if len(variations) == 8 {
				answer = variations[0]
				break substitutions
			}
		}
	}

	fmt.Println(answer)
}
