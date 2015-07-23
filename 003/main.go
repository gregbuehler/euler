package main

import "fmt"
import "math"

const target int64 = 600851475143

func isPrime(n int64) bool {
	var i int64 = 2
	for ; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var i int64 = int64(math.Sqrt(float64(target)))
	for ; i > 1; i-- {
		if target%i == 0 && isPrime(i) {
			fmt.Println("\nAnswer: ", i)
			break
		}
	}
}
