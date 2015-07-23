package main

import "fmt"

func sumOfSquares(n int) int {
	r := 0

	for i := 1; i <= n; i++ {
		r += (i * i)
	}

	return r
}

func squareOfSums(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		r += i
	}

	return (r * r)
}

func main() {
	x := 100
	r := squareOfSums(x) - sumOfSquares(x)

	fmt.Println("\nAnswer: ", r)
}
