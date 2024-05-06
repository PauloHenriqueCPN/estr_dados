package main

import (
	"fmt"
)

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var F1, F2 int
		fmt.Scan(&F1, &F2)
		fmt.Println(mdc(F1, F2))
	}
}
