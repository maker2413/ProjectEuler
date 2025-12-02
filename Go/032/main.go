package main

import (
	"fmt"
	"strconv"
	"strings"
)

func confirmPandigital(str string) bool {
	digits := make(map[string]int)

	for s := range str {
		if string(str[s]) == "0" {
			return false
		}

		digits[string(str[s])]++
	}

	return len(digits) == 9
}

func sumPandigital() int {
	var sum int
	pandigital := make(map[string]int)

	for a := 1; a <= 98; a++ {
		for b := 1; b <= 4321; b++ {
			strA := strconv.Itoa(a)
			strB := strconv.Itoa(b)
			strProd := strconv.Itoa(a * b)

			if len(strA+strB+strProd) == 9 {
				fullString := strings.Builder{}
				if len(strA) < len(strB) {
					fullString.WriteString(strA)
					fullString.WriteString(strB)
					fullString.WriteString(strProd)
				} else {
					fullString.WriteString(strB)
					fullString.WriteString(strA)
					fullString.WriteString(strProd)
				}

				if confirmPandigital(fullString.String()) {
					if pandigital[strProd] == 0 {
						sum += a * b
						pandigital[strProd]++
					}
				}
			}
		}
	}

	return sum
}

func main() {
	answer := sumPandigital()

	fmt.Println("Answer:", answer)
}
