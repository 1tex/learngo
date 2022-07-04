package queue

// An FIFO queue.
type Queue []interface{}

// Pushs the element into the queue.
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

// Pops
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
