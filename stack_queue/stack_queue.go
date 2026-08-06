package stackqueue

// Implement Queue using Stacks
// ----------------------------
// Implementasi queue menggunakan dua stack.
// (LeetCode #232)
//
// Operasi: Push, Pop, Peek, Empty
//
// Contoh:
//
//	q.Push(1); q.Push(2)
//	q.Peek() -> 1
//	q.Pop()  -> 1
//	q.Empty() -> false
type MyQueue struct {
	inbox  []int
	outbox []int
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inbox = append(q.inbox, x)
}

func (q *MyQueue) transfer() {
	if len(q.outbox) == 0 {
		for len(q.inbox) > 0 {
			n := len(q.inbox)
			q.outbox = append(q.outbox, q.inbox[n-1])
			q.inbox = q.inbox[:n-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	n := len(q.outbox)
	val := q.outbox[n-1]
	q.outbox = q.outbox[:n-1]
	return val
}

func (q *MyQueue) Peek() int {
	q.transfer()
	return q.outbox[len(q.outbox)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inbox) == 0 && len(q.outbox) == 0
}
