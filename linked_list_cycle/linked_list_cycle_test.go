package linkedlistcycle

import "testing"

func TestHasCycleTrue(t *testing.T) {
	// 1 -> 2 -> 3 -> back to 2
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	n3.Next = n2 // cycle
	if !hasCycle(n1) {
		t.Error("should have cycle")
	}
}

func TestHasCycleFalse(t *testing.T) {
	// 1 -> 2 -> 3 -> nil
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	if hasCycle(n1) {
		t.Error("should not have cycle")
	}
}

func TestHasCycleSingleNodeNoCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	if hasCycle(n1) {
		t.Error("single node should not have cycle")
	}
}

func TestHasCycleSingleNodeSelfCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n1.Next = n1 // self cycle
	if !hasCycle(n1) {
		t.Error("self-cycle should be true")
	}
}

func TestHasCycleEmpty(t *testing.T) {
	if hasCycle(nil) {
		t.Error("nil should not have cycle")
	}
}

func TestHasCycleCycleAtEnd(t *testing.T) {
	// 1 -> 2 -> 3 -> 3 (self-loop at tail)
	n3 := &ListNode{Val: 3}
	n3.Next = n3
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	if !hasCycle(n1) {
		t.Error("should have cycle")
	}
}
