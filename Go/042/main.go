package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func numericValue(s string) int {
	var value int

	for i := 0; i < len(s); i++ {
		value += int(s[i] - 64)
	}

	return value
}

func main() {
	var answer int
	triangleNumbers := make(map[int]int)

	for i := 1.0; i <= 25; i++ {
		triangleNumbers[int((.5*i)*(i+1.0))] = 1
	}

	f, _ := ioutil.ReadFile("words.txt")

	str := string(f)
	str = strings.ReplaceAll(str, "\"", "")

	words := strings.Split(str, ",")
	sort.Strings(words)

	for j := 0; j < len(words); j++ {
		if triangleNumbers[numericValue(words[j])] == 1 {
			answer++
		}
	}

	fmt.Println(answer)
}
