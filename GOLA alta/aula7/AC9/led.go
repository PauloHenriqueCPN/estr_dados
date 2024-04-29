package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	leds := map[rune]int{
		'0': 6,
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
	}

	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		total := 0
		for _, num := range text {
			total += leds[num]
		}
		fmt.Printf("%d leds\n", total)
	}
}
