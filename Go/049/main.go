package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	increaseBy = 3330
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

func swap(str string, i, j int) string {
	b := []byte(str)

	tmp := b[i]
	b[i] = b[j]
	b[j] = tmp

	return string(b)
}

func permutations(str string, index int, p *[]string) {
	if index == len(str)-1 {
		*p = append(*p, str)
	}

	for i := index; i < len(str); i++ {
		str = swap(str, index, i)
		permutations(str, index+1, p)
		str = swap(str, index, i)
	}
}

func main() {
	var answer string

	for i := 1001; i < 10000-(increaseBy*2); i += 2 {
		if isPrime(i) {
			if isPrime(i+increaseBy) && isPrime(i+(increaseBy*2)) {
				var perms []string
				s := strconv.Itoa(i)

				permutations(string(s[:3]), 0, &perms)

				for p := 0; p < len(perms); p++ {
					n, _ := strconv.Atoi(string(perms[p] + string(s[len(s)-1:])))
					if n == i+increaseBy {
						answer = strconv.Itoa(i)
						answer += strconv.Itoa(i + increaseBy)
						answer += strconv.Itoa(i + (increaseBy * 2))
					}
				}
			}
		}
	}

	fmt.Println(answer)
}
