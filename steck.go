package main

import "fmt"

type Stack struct {
	elements []int
}

// Push добавляет элемент в стек
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop()) // Вывод: 3, true
	fmt.Println(stack.Pop()) // Вывод: 2, true
	fmt.Println(stack.Pop()) // Вывод: 1, true
	fmt.Println(stack.Pop()) // Вывод: 0, false
}
