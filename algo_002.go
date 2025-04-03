package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(8558))
	// []rune(fmt.Sprint(num))
}

func isPalindrome(num int) bool {
	directOrder := strconv.Itoa(num)

	l := len(directOrder) - 1
	j := l
	for i := 0; i <= j; i, j = i+1, j-1 {
		if directOrder[i] == directOrder[j] {
			continue
		}
	}
	return true
}
