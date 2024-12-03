package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	digits     = "0123456789"
	permNumber = 1000000
)

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

func sortedPermutations(str string) []string {
	var perms []string

	permutations(str, 0, &perms)

	return perms
}

func main() {
	var answer string

	perms := sortedPermutations(digits)

	sort.Slice(perms, func(i, j int) bool {
		m, _ := strconv.Atoi(perms[i])
		n, _ := strconv.Atoi(perms[j])

		return m < n
	})

	answer = perms[permNumber-1]

	fmt.Println(answer)
}
