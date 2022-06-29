package main

import (
	"fmt"
	"math"
)

func main() {
	is := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range is {

		mod := math.Mod(float64(n), 2)

		if mod > 0 {
			fmt.Printf("%d is odd\n", n)
		} else {
			fmt.Printf("%d is even\n", n)
		}
	}
}
