package topkfrequent

import (
	"sort"
	"testing"
)

// sortedResult: normalize urutan sebelum compare karena
// elemen dengan frekuensi sama bisa muncul dalam urutan apapun.
func sortedResult(s []int) []int {
	cp := make([]int, len(s))
	copy(cp, s)
	sort.Ints(cp)
	return cp
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int // sorted ascending untuk perbandingan
	}{
		{"k=2 basic", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"k=1 single", []int{1}, 1, []int{1}},
		{"k=2 tie", []int{1, 2}, 2, []int{1, 2}},
		{"k=1 clear winner", []int{4, 4, 4, 2, 2, 3}, 1, []int{4}},
		{"all same", []int{5, 5, 5}, 1, []int{5}},
		{"k=3", []int{1, 1, 2, 2, 3, 3, 4}, 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedResult(topKFrequent(tt.nums, tt.k))
			want := sortedResult(tt.want)
			if len(got) != len(want) {
				t.Fatalf("topKFrequent(%v, %d) len=%d, want len=%d", tt.nums, tt.k, len(got), len(want))
			}
			for i := range got {
				if got[i] != want[i] {
					t.Errorf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, want)
					return
				}
			}
		})
	}
}
