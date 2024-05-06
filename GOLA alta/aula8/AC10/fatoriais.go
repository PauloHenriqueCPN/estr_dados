package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	return fact
}

func main() {
	var M, N int64
	for {
		_, err := fmt.Scan(&M, &N)
		if err != nil {
			break
		}
		factM := factorial(M)
		factN := factorial(N)
		sum := big.NewInt(0).Add(factM, factN)
		fmt.Println(sum)
	}
}
