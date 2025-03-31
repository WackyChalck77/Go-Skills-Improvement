package main

import "fmt"

func main() {
	fmt.Println(remove("Poluchil_po_UE-ball", 12))
	fmt.Println(remove("1234567889", 7))

}

func remove(str string, index int) string {
	runeSlice := []rune(str)
	var resultSlice []rune
	for i, r := range runeSlice {
		if i != index {
			resultSlice = append(resultSlice, r)
			//fmt.Println(string(resultSlice))
		}
	}
	return string(resultSlice)
}
