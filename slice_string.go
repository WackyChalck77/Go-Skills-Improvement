package main

import (
	"fmt"
)

func main() {
	str := "Hello!!!Привет"
	byteSlice := []byte(str)
	fmt.Println(byteSlice)

	byteSlice2 := []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	str2 := string(byteSlice2)
	fmt.Println(str2)

}
