package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, X int
	fmt.Scan(&N)

	numbers := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scan(&X)
		numbers[X]++
	}

	keys := make([]int, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d aparece %d vez(es)\n", k, numbers[k])
	}
}
