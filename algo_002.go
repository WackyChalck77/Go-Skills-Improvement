package main

import (
	"fmt"
	//"strconv"
)

func main() {
	//fmt.Println(isPalindrome(8886666888))
	// []rune(fmt.Sprint(num))
	isIt := isPalindrome(5)
	fmt.Println(isIt)
}

func isPalindrome(num int) bool {

	var runeSlice1 []rune
	var rune1 rune
	var digit int

	for num > 0 {
		digit = num % 10
		rune1 = rune(digit)
		runeSlice1 = append(runeSlice1, rune1)
		num = num / 10
	}
	l := len(runeSlice1) - 1
	j := l

	for i := 0; i < j; i++ {
		if runeSlice1[i] != runeSlice1[j] {
			return false
		}
		j--
	}
	return true
}
