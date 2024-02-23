package main

import (
	"fmt"
)

func calculaMedia(x ...float32) float32 {
	soma := float32(0)
	for _, sn := range x {
		soma += sn
		
	}
	
	sa := len(x)
	media := float32(soma) / float32(sa)
	return media
}

func verificaParidade(x int) {
	if x % 2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}
}

func minhaPotencia(a , b int) int {
	base := 1
	for e := 1; e <= b; e++ {
		base *= a
	}
	
	return base
}

func converteCelsiusParaFahrenheit(c int) {
	var f int
	f = ((c / 5) * 9) + 32
	fmt.Println(f)
}

func main() {
	media := calculaMedia(1, 2.5, 3, 4.5, 5)
	fmt.Println(media)

	verificaParidade(10)

	base := minhaPotencia(2,4)
	fmt.Println(base)

	converteCelsiusParaFahrenheit(20)
}
