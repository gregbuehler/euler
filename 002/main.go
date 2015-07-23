package main

import "fmt"

func main() {
	r := 0

	j, k := 1, 0
	for r < 4e6 {
		f := j + k
		if f%2 == 0 {
			r += f
		}

		k = j
		j = f
	}

	fmt.Println("\nAnswer: ", r)
}
