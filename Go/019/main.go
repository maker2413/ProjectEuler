package main

import "fmt"

func checkSunday(n int) bool {
	if n%7 == 0 {
		return true
	}

	return false
}

func main() {
	var sundays int

	day := 1   // January 1st 1900
	day += 365 // January 1st 1901

	for year := 1901; year <= 2000; year++ {
		if checkSunday(day) {
			sundays++
		}
		day += 31 // January
		if checkSunday(day) {
			sundays++
		}
		if year%4 == 0 {
			day += 29 // February
		} else {
			day += 28 // February
		}
		if checkSunday(day) {
			sundays++
		}
		day += 31 // March
		if checkSunday(day) {
			sundays++
		}
		day += 30 // April
		if checkSunday(day) {
			sundays++
		}
		day += 31 // May
		if checkSunday(day) {
			sundays++
		}
		day += 30 // June
		if checkSunday(day) {
			sundays++
		}
		day += 31 // July
		if checkSunday(day) {
			sundays++
		}
		day += 31 // August
		if checkSunday(day) {
			sundays++
		}
		day += 30 // September
		if checkSunday(day) {
			sundays++
		}
		day += 31 // October
		if checkSunday(day) {
			sundays++
		}
		day += 30 // November
		if checkSunday(day) {
			sundays++
		}
		day += 31 // December
	}

	fmt.Println(sundays)
}
