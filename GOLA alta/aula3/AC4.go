package main

import "fmt"


func torreHanoi(numDisc int, origem, destino, auxiliar string) {
	if numDisc == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", origem, destino)
		return
	}

	torreHanoi(numDisc-1, origem, auxiliar, destino)
	fmt.Printf("Mova o disco %d de %s para %s\n", numDisc, origem, destino)
	torreHanoi(numDisc-1, auxiliar, destino, origem)

}


func encontraPar(array []int, alvo int) (int, int) {
	inicio := 0
	fim := len(array) - 1

	for inicio < fim {
		soma := array[inicio] + array[fim]

		if soma == alvo {
			return array[inicio], array[fim]
		} else if soma < alvo {
			inicio++
		} else {
			fim--
		}
	}
	
	return -1, -1
	
}


func main() {
	torreHanoi(4, "torre1", "torre2", "torre3")


	num1, num2 := encontraPar([]int{1,2,3,4,5,6,7,8,9}, 10)

	if num1 != -1 && num2 != -1 {
		fmt.Printf("\nPar encontrado: %d, %d\n", num1, num2)
	} else {
		fmt.Println("\nNenhum par encontrado: -1, -1")
	}

}
