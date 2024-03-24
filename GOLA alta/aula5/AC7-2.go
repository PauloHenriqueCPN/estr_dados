package main

import "fmt"

func main() {

	var numPart int
	var garrafa, cuia float64
	var nomesL []string

	part := 0

	fmt.Scan(&numPart, &garrafa, &cuia)

	for x := 0; x < numPart; x++ {
		var nome string
		fmt.Scan(&nome)
		nomesL = append(nomesL, nome)
	}

	for garrafa-cuia > 0 {
		garrafa -= cuia
		part++
	}

	part = part % numPart
	fmt.Printf("%s %.1f\n", nomesL[part], garrafa)

}
