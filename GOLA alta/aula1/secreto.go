package main

import "fmt"

func torreHanoi() {
	var a1, a2, a3, a4, a5, a6, x string
	var eArgola, ePoste int

	a1 = "argola1"
	a2 = "argola2"
	a3 = "argola3"
	a4 = "argola4"
	a5 = "argola5"
	a6 = "argola6"

	eArgola = 0
	ePoste = 0
	
	fmt.Println(" poste1      poste2      poste3")
	fmt.Println("",a1,"\n",a2,"\n",a3,"\n",a4,"\n",a5,"\n",a6)
	fmt.Println("Escolha uma argola para mover(apenas o número):")
	fmt.Scanln(&eArgola)
	fmt.Println("Escolha em que poste quer colocar esta argola(apenas o número)")
	fmt.Scanln(&ePoste)

	switch eArgola {
	case 1:
		x = a1
	case 2:
		x = a2
	case 3:
		x = a3
	case 4:
		x = a4
	case 5:
		x = a5
	case 6:
		x = a6
	}

	fmt.Println(x)

	
}

func meteoritos() {
	var x1, x2, y1, y2, x, y, n int
	num := make([]int, 0)
	i := 0
	for {
		fmt.Println("x1, x2, x3, x4")
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 { break }
		fmt.Println("n")
		fmt.Scanln(&n)
		num = append(num, 0)
		for j := 1; j <= n; j++ {
			fmt.Println("x e y")
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				num[i]++
			}
		}
		i++
	}
	for j, num_meteoritos := range num {
		fmt.Println("Teste", j + 1)
		fmt.Println(num_meteoritos)
		fmt.Println("")
	}
}

func main() {
	torreHanoi()
	// meteoritos()
}
