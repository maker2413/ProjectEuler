package main

import (
	"fmt"
	"strings"
)

const (
	to = 1000
)

var words = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func convertToWords(n int, result *strings.Builder) string {
	if n < 0 {
		result.WriteString("negative")
		n = -n
	}

	if len(result.String()) > 0 {
		if n == 0 {
			return result.String()
		}
		result.WriteString(" ")
	}

	switch {
	case n > 9999:
		result.WriteString("Numbers this large are currently not supported")
		n = 0
	case n >= 1000:
		m := firstDigit(n)
		result.WriteString(words[m])
		result.WriteString(" ")
		result.WriteString("thousand")
		convertToWords(n-(m*1000), result)
	case n >= 100:
		m := firstDigit(n)
		result.WriteString(words[m])
		result.WriteString(" ")
		result.WriteString("hundred")
		if n-(m*100) != 0 {
			result.WriteString(" ")
			result.WriteString("and")
		}
		convertToWords(n-(m*100), result)
	case n >= 90:
		result.WriteString(words[90])
		convertToWords(n-90, result)
	case n >= 80:
		result.WriteString(words[80])
		convertToWords(n-80, result)
	case n >= 70:
		result.WriteString(words[70])
		convertToWords(n-70, result)
	case n >= 60:
		result.WriteString(words[60])
		convertToWords(n-60, result)
	case n >= 50:
		result.WriteString(words[50])
		convertToWords(n-50, result)
	case n >= 40:
		result.WriteString(words[40])
		convertToWords(n-40, result)
	case n >= 30:
		result.WriteString(words[30])
		convertToWords(n-30, result)
	case n >= 20:
		result.WriteString(words[20])
		convertToWords(n-20, result)
	default:
		result.WriteString(words[n])
	}

	return result.String()
}

func firstDigit(n int) int {
	for n >= 10 {
		n /= 10
	}

	return n
}

func characterCount(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")

	return len(s)
}

func main() {
	var answer int

	for i := 1; i <= to; i++ {
		var result strings.Builder

		answer += characterCount(convertToWords(i, &result))
	}

	fmt.Println(answer)
}
