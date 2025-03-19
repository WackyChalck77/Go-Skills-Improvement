package main

import "fmt"

type Queue struct {
	elements []int
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	element := q.elements[0]
	q.elements = q.elements[1:] //выбор со второго элемента
	return element, true
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // Вывод: 1, true
	fmt.Println(queue.Dequeue()) // Вывод: 2, true
	fmt.Println(queue.Dequeue()) // Вывод: 3, true
	fmt.Println(queue.Dequeue()) // Вывод: 0, false
}
