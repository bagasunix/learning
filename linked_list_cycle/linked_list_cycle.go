package linkedlistcycle

// Linked List Cycle (Floyd's Algorithm)
// -------------------------------------
// Cek apakah linked list memiliki cycle.
// (LeetCode #141)
//
// Contoh:
//
//	1 -> 2 -> 3 -> 2 (cycle)  -> true
//	1 -> 2 -> 3               -> false
//
// Hint: Gunakan two-pointer (slow & fast), O(1) space.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
