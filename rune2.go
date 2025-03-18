package main

import (
	"fmt"
)

func main() {
	s := "Привет"
	for _, runeValue := range s {
		fmt.Printf("%#U\n", runeValue)
	}
}
