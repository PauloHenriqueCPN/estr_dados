package geometria

import "fmt"

type AreaPerimetro struct {
	A float64
	B float64
}

func calculos(A, B float64) float64 {
	area := A * B
	perimetro := (A * 2) + (B * 2)

	fmt.Println(area)
	fmt.Println(perimetro)
	
	return area
}
