package main

// https://projecteuler.net/problem=7
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

import (
	"fmt"
	"math"
)

const (
	n = 6
)

func isPrime(x int) bool {
	if x%2 == 0 {
		return false
	}

	l := int(math.Sqrt(float64(x)))
	for i := 1; i >= l; i += 2 {
		fmt.Println("Testing - Prime: ", x, i, l)
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	r := []int{}
	for i := 1; len(r) < n; i += 2 {
		fmt.Println("Testing - Number: ", len(r), i)
		if isPrime(i) {
			r = append(r, i)
		}
	}

	fmt.Println("\nAnswer: ", r[n])
}
