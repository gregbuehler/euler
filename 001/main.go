package main

import "fmt"

var (
	target      = 1000
	multiplesof = []int{3, 5}
)

func main() {

	r := 0
	for i := 1; i < target; i++ {
		for _, j := range multiplesof {
			if i%j == 0 {
				r += i
				break
			}
		}
	}

	fmt.Println(r)
}
