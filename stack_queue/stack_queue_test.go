package stackqueue

import "testing"

func TestQueueBasic(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 {
		t.Errorf("Peek got %d, want 1", q.Peek())
	}
	if q.Pop() != 1 {
		t.Errorf("Pop got %d, want 1", q.Pop())
	}
	if q.Empty() {
		t.Error("should not be empty")
	}
}

func TestQueuePopAll(t *testing.T) {
	q := NewQueue()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	if q.Pop() != 10 {
		t.Error("first pop should be 10")
	}
	if q.Pop() != 20 {
		t.Error("second pop should be 20")
	}
	if q.Pop() != 30 {
		t.Error("third pop should be 30")
	}
	if !q.Empty() {
		t.Error("should be empty after popping all")
	}
}

func TestQueueInterleaved(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	if q.Pop() != 1 {
		t.Error("pop should be 1")
	}
	q.Push(3)
	if q.Peek() != 2 {
		t.Errorf("peek got %d, want 2", q.Peek())
	}
	if q.Pop() != 2 {
		t.Error("pop should be 2")
	}
	if q.Pop() != 3 {
		t.Error("pop should be 3")
	}
	if !q.Empty() {
		t.Error("should be empty")
	}
}

func TestQueueEmpty(t *testing.T) {
	q := NewQueue()
	if !q.Empty() {
		t.Error("new queue should be empty")
	}
}

func TestQueueSinglePushPop(t *testing.T) {
	q := NewQueue()
	q.Push(42)
	if q.Empty() {
		t.Error("should not be empty after push")
	}
	if q.Pop() != 42 {
		t.Error("pop should return 42")
	}
	if !q.Empty() {
		t.Error("should be empty after pop")
	}
}
