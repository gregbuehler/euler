package main

import "fmt"

var (
	factor = 20
)

func main() {

	solved := false
	r := factor

	for i := factor; solved != true; i += factor {
		for j := factor; j > 0; j-- {
			if i%j != 0 {
				break
			}
			if j == 1 {
				solved = true
				break
			}
		}

		r = i
	}

	fmt.Println("\nAnswer: ", r)
}
