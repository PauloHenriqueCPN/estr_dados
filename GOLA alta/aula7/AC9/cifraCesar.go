package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func decrypt(c rune, shift int) rune {
	if c >= 'A' && c <= 'Z' {
		return ((c-'A'+26-int32(shift))%26) + 'A'
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		text := scanner.Text()
		scanner.Scan()
		shift, _ := strconv.Atoi(scanner.Text())
		fmt.Println(strings.Map(func(c rune) rune { return decrypt(c, shift) }, text))
	}
}
