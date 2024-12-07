package main

import (
	"fmt"
	"strconv"
)

func isPandigital(str string) bool {
	digits := make(map[string]int)

	for s := 0; s < len(str); s++ {
		digits[string(str[s])]++
	}

	if len(digits) == 10 {
		return true
	}

	return false
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

func isDivisibile(str string) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17}

	for i := 0; i < len(primes); i++ {
		n, _ := strconv.Atoi(str[i+1 : len(str)-6+i])

		if n%primes[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	var answer int
	var perms []string

	permutations("0123456789", 0, &perms)

	for i := 0; i < len(perms); i++ {
		if isDivisibile(perms[i]) {
			n, _ := strconv.Atoi(perms[i])

			answer += n
		}
	}

	fmt.Println(answer)
}
