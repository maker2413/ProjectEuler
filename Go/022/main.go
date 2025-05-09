package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func numericValue(s string) int {
	var value int

	for i := range s {
		value += int(s[i] - 64)
	}

	return value
}

func main() {
	var answer int

	f, _ := os.ReadFile("names.txt")

	str := string(f)
	str = strings.ReplaceAll(str, "\"", "")

	names := strings.Split(str, ",")
	sort.Strings(names)

	for i := range names {
		answer += (i + 1) * numericValue(names[i])
	}

	fmt.Println(answer)
}
