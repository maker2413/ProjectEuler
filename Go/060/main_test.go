package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{337, true},
		{373, true},
		{4294967291, true},
	}

	for _, test := range tests {
		result := isPrime(test.number)
		assert.Equal(t, test.expected, result,
			"got: %t, want: %t", result, test.expected)
	}
}

func TestPrimeSieve(t *testing.T) {
	primeLimit := 100
	primeList := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	result := primeSieve(primeLimit)
	assert.Equal(t, primeList, result,
		"got: %d, want: %d", result, primeList)
}

func TestConcatenateInts(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{3, 7, 37},
		{3, 109, 3109},
		{3, 673, 3673},
		{7, 3, 73},
		{7, 109, 7109},
		{7, 673, 7673},
		{109, 673, 109673},
		{109, 7, 1097},
		{109, 3, 1093},
		{673, 7, 6737},
		{673, 3, 6733},
		{673, 109, 673109},
	}

	for _, test := range tests {
		result := concatenateInts(test.a, test.b)
		assert.Equal(t, test.expected, result,
			"got: %d, want: %d", result, test.expected)
	}
}

func TestIsPrimePair(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected bool
	}{
		{3, 7, true},
		{3, 109, true},
		{3, 673, true},
		{3, 37, true},
		{3, 19, false},
	}

	for _, test := range tests {
		result := isPrimePair(test.a, test.b)
		assert.Equal(t, test.expected, result,
			"got: %t, want: %t", result, test.expected)
	}
}

func TestSumSet(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{3}, 3},
		{[]int{3, 7, 19, 109}, 138},
	}

	for _, test := range tests {
		result := sumSet(test.nums)
		assert.Equal(t, test.expected, result,
			"got: %d, want: %d", result, test.expected)
	}
}
