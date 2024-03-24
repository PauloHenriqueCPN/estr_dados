package main

import "fmt"

func main() {
    var codP1, numP1, codP2, numP2 int
    var valorUnitarioPeca1, valorUnitarioPeca2 float64

    fmt.Scan(&codP1, &numP1, &valorUnitarioPeca1)

    fmt.Scan(&codP2, &numP2, &valorUnitarioPeca2)

    valorTotal := (float64(numP1) * valorUnitarioPeca1) + (float64(numP2) * valorUnitarioPeca2)

    fmt.Printf("Valor a pagar: R$ %.2f\n", valorTotal)
}
