package main

import (
	"fmt"
	"slices"
	"strconv"
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

func main() {
	var answer int
	i := 1

	for answer == 0 {
		if len(strconv.Itoa(i)) == len(strconv.Itoa(i*6)) {
			var perms []string
			permutations(strconv.Itoa(i), 0, &perms)
			for j := 2; j <= 6; j++ {
				if !slices.Contains(perms, strconv.Itoa(i*j)) {
					break
				}

				if j == 6 {
					answer = i
				}
			}
		}
		i++
	}

	fmt.Println("Answer:", answer)
}
