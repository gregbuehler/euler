package main

import (
	"fmt"
	"strconv"
)

var (
	factor = 1000
)

func isPalindrone(s string) bool {
	z := int(len(s) / 2)
	for i := 0; i < z; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	r := 0

	for j := 1; j <= factor; j++ {
		for k := 1; k <= factor; k++ {
			t := j * k

			if t > r && isPalindrone(strconv.Itoa(t)) {
				r = t
			}
		}
	}

	fmt.Println("\nAnswer: ", r)
}
