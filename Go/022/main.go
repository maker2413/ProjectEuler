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

	f, _ := ioutil.ReadFile("names.txt")

	str := string(f)
	str = strings.ReplaceAll(str, "\"", "")

	names := strings.Split(str, ",")
	sort.Strings(names)

	for i := 0; i < len(names); i++ {
		answer += (i + 1) * numericValue(names[i])
	}

	fmt.Println(answer)
}
