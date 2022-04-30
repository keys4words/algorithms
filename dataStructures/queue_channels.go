package main

import "fmt"

type QueueChan struct {
	items chan int
}

func (q *QueueChan) Enqueue(item int) {
	q.items <- item
}

func (q *QueueChan) Dequeue() int {
	return <-q.items
}

func main() {
	q := QueueChan{items: make(chan int, 16)}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
}
