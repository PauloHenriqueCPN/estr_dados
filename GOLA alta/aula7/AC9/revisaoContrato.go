package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "0 0" {
			break
		}
		parts := strings.Split(line, " ")
		digit := parts[0]
		number := parts[1]
		result := strings.Replace(number, digit, "", -1)
		if result == "" {
			fmt.Println("0")
		} else {
			result = strings.TrimLeft(result, "0")
			if result == "" {
				fmt.Println("0")
			} else {
				fmt.Println(result)
			}
		}
	}
}