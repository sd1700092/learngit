package queue

type MyQueue []int

func (q *MyQueue) Push(v int) {
	*q = append(*q, v)
}

func (q *MyQueue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *MyQueue) IsEmpty() bool {
	return len(*q) == 0
}
