package stackqueue

// =============================================================================
// IMPLEMENT QUEUE USING TWO STACKS
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Stack, amortized analysis, data structure design
// Complexity: Push O(1), Pop/Peek amortized O(1), worst case O(n)
// =============================================================================
//
// SOAL:
//   Implementasikan Queue (FIFO) menggunakan dua Stack (LIFO).
//   (LeetCode #232)
//
//   Operasi yang harus didukung:
//   - Push(x) : tambah elemen ke belakang antrian
//   - Pop()   : hapus dan return elemen paling depan
//   - Peek()  : return elemen paling depan tanpa hapus
//   - Empty() : return true jika antrian kosong
//
// CONTOH:
//   q := NewQueue()
//   q.Push(1); q.Push(2); q.Push(3)
//   q.Peek()  -> 1   (elemen pertama masuk)
//   q.Pop()   -> 1
//   q.Pop()   -> 2
//   q.Empty() -> false
//   q.Pop()   -> 3
//   q.Empty() -> true
//
// EDGE CASES yang harus disebut:
//   - Pop/Peek pada queue kosong -> undefined (soal asumsi tidak terjadi)
//   - Push banyak lalu Pop bergantian
//
// KENAPA PAKAI DUA STACK?
//   Stack = LIFO (Last In First Out)
//   Queue = FIFO (First In First Out)
//   Dua stack bisa "membalik" urutan:
//   Stack 1 (inbox):  [1,2,3]  (3 di atas)
//   Transfer ke Stack 2 (outbox): [3,2,1]  (1 di atas = urutan queue!)
//
// CARA KERJA (inbox/outbox):
//   Push selalu ke inbox.
//   Pop/Peek: kalau outbox kosong, transfer semua dari inbox ke outbox dulu.
//   Kenapa transfer HANYA KALAU outbox kosong? -> amortized O(1)
//   Setiap elemen pindah dari inbox ke outbox tepat SATU KALI seumur hidupnya.
//
//   Push(1), Push(2), Push(3):  inbox=[1,2,3], outbox=[]
//   Pop(): outbox kosong -> transfer: inbox=[], outbox=[3,2,1]
//          pop dari outbox: return 1, outbox=[3,2]
//   Pop(): outbox tidak kosong -> pop langsung: return 2, outbox=[3]
//   Push(4):                    inbox=[4], outbox=[3]
//   Pop(): outbox tidak kosong -> pop: return 3, outbox=[]
//   Pop(): outbox kosong -> transfer: inbox=[], outbox=[4]
//          pop: return 4 ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Inbox untuk push, outbox untuk pop/peek.
//    Transfer dari inbox ke outbox HANYA ketika outbox kosong.
//    Setiap elemen di-transfer tepat sekali -> amortized O(1) per operasi."
//
// JEBAKAN UMUM:
//   - Transfer setiap kali Pop -> O(n) setiap Pop, total O(n²)
//   - Transfer hanya kalau outbox kosong -> amortized O(1) ← BENAR
// =============================================================================

type MyQueue struct {
	inbox  []int // stack untuk push
	outbox []int // stack untuk pop/peek
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

// Push: selalu ke inbox — O(1)
func (q *MyQueue) Push(x int) {
	q.inbox = append(q.inbox, x)
}

// transfer: pindahkan inbox ke outbox (hanya kalau outbox kosong)
func (q *MyQueue) transfer() {
	if len(q.outbox) == 0 {
		for len(q.inbox) > 0 {
			n := len(q.inbox)
			q.outbox = append(q.outbox, q.inbox[n-1])
			q.inbox = q.inbox[:n-1]
		}
	}
}

// Pop: hapus dan return elemen paling depan — amortized O(1)
func (q *MyQueue) Pop() int {
	q.transfer()
	n := len(q.outbox)
	val := q.outbox[n-1]
	q.outbox = q.outbox[:n-1]
	return val
}

// Peek: lihat elemen paling depan tanpa hapus — amortized O(1)
func (q *MyQueue) Peek() int {
	q.transfer()
	return q.outbox[len(q.outbox)-1]
}

// Empty: cek apakah queue kosong — O(1)
func (q *MyQueue) Empty() bool {
	return len(q.inbox) == 0 && len(q.outbox) == 0
}
