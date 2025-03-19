package main

import "fmt"

func main() {
	// Создание и инициализация массива
	arr := []int{1, 2, 3}

	// Добавление элемента в массив
	arr = append(arr, 4)

	// Печать всех элементов массива
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println() // Вывод: 1 2 3 4
}
