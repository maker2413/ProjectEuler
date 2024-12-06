package main

import (
	"fmt"
	"strconv"
	"strings"
)

func confirmPandigital(str string) bool {
	digits := make(map[string]int)

	for s := 0; s < len(str); s++ {
		if string(str[s]) == "0" {
			return false
		}

		digits[string(str[s])]++
	}

	if len(digits) == 9 {
		return true
	}

	return false
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
						fmt.Println(a, b, a*b)
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
	//answer := confirmPandigital("437691258")

	fmt.Println(answer)
	//fmt.Println(confirmPandigital("123456780"))
}
