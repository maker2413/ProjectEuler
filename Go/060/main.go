package main

import (
	"fmt"
	"math"
)

const (
	primeLimit = 20000
)

func primeSieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	primes := []int{2}

	n := (limit - 3) / 2
	isComposite := make([]bool, n+1)

	sqrtLimit := int(math.Sqrt(float64(limit)))

	for i := 0; (2*i + 3) <= sqrtLimit; i++ {
		if !isComposite[i] {
			p := 2*i + 3

			startIndex := (p*p - 3) / 2

			for j := startIndex; j <= n; j += p {
				isComposite[j] = true
			}
		}
	}

	for i := 0; i <= n; i++ {
		if !isComposite[i] {
			primes = append(primes, 2*i+3)
		}
	}

	return primes
}

func concatenateInts(a, b int) int {
	temp := b
	for temp > 0 {
		a *= 10
		temp /= 10
	}

	return a + b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func isPrimePair(a, b int) bool {
	return isPrime(concatenateInts(a, b)) && isPrime(concatenateInts(b, a))
}

func checkSet(nums []int) bool {
	for a := range len(nums) - 1 {
		for b := a + 1; b < len(nums); b++ {
			if !isPrimePair(nums[a], nums[b]) {
				return false
			}
		}
	}

	return true
}

func findSet(nums []int) []int {
	set := []int{}

	for a := 1; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			if !isPrimePair(nums[a], nums[b]) {
				continue
			}

			for c := b + 1; c < len(nums)-1; c++ {
				if !isPrimePair(nums[a], nums[c]) {
					continue
				}
				for d := c + 1; d < len(nums)-1; d++ {
					if !isPrimePair(nums[a], nums[d]) {
						continue
					}

					set = []int{nums[0], nums[a], nums[b], nums[c], nums[d]}
					if checkSet(set) {
						return set
					}
				}
			}
		}
	}

	return nil
}

func sumSet(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func main() {
	answer := primeLimit

	primes := primeSieve(primeLimit)

	for a := range len(primes) - 4 {
		primePairs := []int{primes[a]}

		for b := a + 1; b < len(primes); b++ {
			if !isPrimePair(primes[a], primes[b]) {
				continue
			}

			primePairs = append(primePairs, primes[b])
		}

		if len(primePairs) > 5 {
			sum := sumSet(findSet(primePairs))

			if sum < answer && sum != 0 {
				answer = sum
			}
		}
	}

	fmt.Println("Answer:", answer)
}
