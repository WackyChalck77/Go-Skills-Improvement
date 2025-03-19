package main

import "fmt"

func main() {
	fmt.Println("Цикл For")
	str := "Пример строки"
	for i := 0; i < len(str); i++ {
		fmt.Printf("Символ на позиции %d: %c\n", i, str[i])
	}

	fmt.Println("Цикл Range")
	str = "Пример строки"
	for pos, char := range str {
		fmt.Printf("Символ на позиции %d: %c\n", pos, char)
	}
}
