package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	for {
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}
		fmt.Println(minMovimentos(x1, y1, x2, y2))
	}
}

func minMovimentos(x1, y1, x2, y2 int) int {
	if x1 == x2 && y1 == y2 {
		return 0
	}
	if x1+y1 == x2+y2 || x1-y1 == x2-y2 || x1 == x2 || y1 == y2 {
		return 1
	}
	return 2
}
