package linkedlistcycle

// =============================================================================
// LINKED LIST CYCLE (Floyd's Tortoise and Hare)
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Two-pointer (slow & fast), linked list
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Cek apakah sebuah linked list memiliki cycle (lingkaran).
//   (LeetCode #141)
//
// CONTOH:
//   1 -> 2 -> 3 -> 4 -> 2 (node 4 menunjuk ke node 2) -> true
//   1 -> 2 -> 3 -> nil                                 -> false
//   1 -> 1 (menunjuk ke dirinya sendiri)               -> true
//   nil                                                -> false
//
// EDGE CASES yang harus disebut:
//   - head = nil           -> false (list kosong)
//   - Satu node, tanpa cycle -> false
//   - Satu node, menunjuk ke dirinya -> true
//
// CARA KERJA Floyd's Algorithm:
//   Bayangkan slow berjalan 1 langkah, fast 2 langkah.
//   Kalau ada cycle, fast pasti akan "mengejar" slow dari belakang
//   dan keduanya akan bertemu di dalam cycle.
//
//   Tanpa cycle:
//   fast akan mencapai nil duluan -> return false
//
//   Dengan cycle (1->2->3->4->2):
//   Step 1: slow=2, fast=3
//   Step 2: slow=3, fast=2  (fast melompat 3->4->2)
//   Step 3: slow=4, fast=4  (fast melompat 2->3->4... eh tunggu)
//   slow=4, fast: 3->4... slow==fast -> true ✓
//
// KENAPA O(1) SPACE?
//   Approach naif: simpan semua node yang sudah dikunjungi di hash set -> O(n) space
//   Floyd's: hanya dua pointer -> O(1) space
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Dua pointer: slow maju 1 langkah, fast maju 2 langkah.
//    Kalau fast atau fast.Next nil -> tidak ada cycle, return false.
//    Kalau slow == fast -> ada cycle, return true.
//    Ini O(n) time, O(1) space karena tidak perlu menyimpan node yang dikunjungi."
//
// JEBAKAN UMUM:
//   - Cek fast != nil saja tidak cukup, harus cek fast.Next != nil juga
//     (karena kita akses fast.Next.Next)
//   - Pakai hash set -> O(n) space, valid tapi bukan optimal
// =============================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle: Floyd's tortoise & hare — O(n) time, O(1) space
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // maju 1
		fast = fast.Next.Next // maju 2
		if slow == fast {
			return true // bertemu -> ada cycle
		}
	}
	return false
}
