package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var C float64
		fmt.Scan(&C)

		dias := 0
		for C > 1.0 {
			C /= 2
			dias++
		}

		fmt.Printf("%d dias\n", dias)
	}
}
