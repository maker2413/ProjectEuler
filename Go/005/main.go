package main

import "fmt"

func main() {
	x, y := 1, 1

	// We can save a huge amount of CPU time by finding the least common multiple
	// of 20 and 19 and then finding multples of that number until we find a number
	// divisible by 18 and multiple that until we find a multiple divisible by 17 etc.
	for i := 20; i > 0; i-- {
		for y%i != 0 {
			y += x
		}
		x = y
	}

	fmt.Println(x)
}
