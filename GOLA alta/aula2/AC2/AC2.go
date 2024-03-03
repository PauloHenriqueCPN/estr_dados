package main

import (
	"fmt"
	"math"
	m "projeto/geometria"
)

func vetor() {
	vet := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, vet := range vet {
		fmt.Println(vet)
	}
}

func nomeInvertido(x string) string {
	runa := []rune(x)
	for i, j := 0, len(runa)-1; i < j; i, j = i+1, j-1 {

		runa[i], runa[j] = runa[j], runa[i]
	}

	fmt.Println(x)
	return string(runa)

}

type Ponto struct {
	x float64
	y float64
}

func (d Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(math.Pow(d.x, 2) + math.Pow(d.y, 2))
}

func main() {
	vetor()
	fmt.Println("----------------------------------------")

	reverso := nomeInvertido("paralelepipedo")
	fmt.Println(reverso)
	fmt.Println("----------------------------------------")

	orig := Ponto{x: 4 , y: 10}
	fmt.Println(orig.DistanciaOrigem())
	fmt.Println("----------------------------------------")

	lado := m.AreaPerimetro{}
	lado.A = 2
	lado.B = 4

	fmt.Println(lado)

}
