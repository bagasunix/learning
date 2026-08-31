package pascaltriangle

import "testing"

// =============================================================================
// TEST — REVERSED PASCAL'S TRIANGLE
// Jalankan: go test ./pascal_triangle/... -v
// =============================================================================

func TestBuildTriangle(t *testing.T) {
	got := buildTriangle(4)
	want := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
	}
	if len(got) != len(want) {
		t.Fatalf("buildTriangle(4) len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("row %d len = %d, want %d", i+1, len(got[i]), len(want[i]))
		}
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("row %d idx %d = %d, want %d", i+1, j, got[i][j], want[i][j])
			}
		}
	}
}

func TestLabelRow(t *testing.T) {
	tests := []struct {
		row  []int
		want []string
	}{
		{[]int{1}, []string{"1"}},
		{[]int{1, 1}, []string{"1", "1"}},
		{[]int{1, 2, 1}, []string{"1", "2gn", "1"}},
		{[]int{1, 3, 3, 1}, []string{"1", "3gj", "3gj", "1"}},
		{[]int{1, 4, 6, 4, 1}, []string{"1", "4gn", "6gn", "4gn", "1"}},
		{[]int{1, 5, 10, 10, 5, 1}, []string{"1", "5gj", "10gn", "10gn", "5gj", "1"}},
	}
	for _, tt := range tests {
		got := labelRow(tt.row)
		if len(got) != len(tt.want) {
			t.Fatalf("labelRow(%v) len = %d, want %d", tt.row, len(got), len(tt.want))
		}
		for i := range tt.want {
			if got[i] != tt.want[i] {
				t.Errorf("labelRow(%v)[%d] = %q, want %q", tt.row, i, got[i], tt.want[i])
			}
		}
	}
}

func TestRowSum(t *testing.T) {
	tests := []struct {
		row  []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 1}, 0},
		{[]int{1, 2, 1}, 2},
		{[]int{1, 3, 3, 1}, 6},
		{[]int{1, 4, 6, 4, 1}, 14},
	}
	for _, tt := range tests {
		if got := rowSum(tt.row); got != tt.want {
			t.Errorf("rowSum(%v) = %d, want %d", tt.row, got, tt.want)
		}
	}
}

func TestReversedPascalTriangle(t *testing.T) {
	results := ReversedPascalTriangle(4)
	if len(results) != 4 {
		t.Fatalf("len = %d, want 4", len(results))
	}

	// urutan dibalik: elemen pertama = baris 4 (terbanyak elemen), terakhir = baris 1
	wantLabels := []string{
		"1 3gj 3gj 1",
		"1 2gn 1",
		"1 1",
		"1",
	}
	wantSums := []int{6, 2, 0, 0}

	for i, r := range results {
		if r.String() != wantLabels[i] {
			t.Errorf("row %d labels = %q, want %q", i, r.String(), wantLabels[i])
		}
		if r.Sum != wantSums[i] {
			t.Errorf("row %d sum = %d, want %d", i, r.Sum, wantSums[i])
		}
	}

	// baris pertama hasil (baris ke-n asli) harus punya elemen terbanyak,
	// baris terakhir (baris ke-1 asli) harus 1 elemen -> ini yang membuatnya "terbalik"
	if len(results[0].Labels) != 4 {
		t.Errorf("results[0] elemen = %d, want 4 (baris terlebar di atas)", len(results[0].Labels))
	}
	if len(results[len(results)-1].Labels) != 1 {
		t.Errorf("results terakhir elemen = %d, want 1 (baris 1 di bawah)", len(results[len(results)-1].Labels))
	}
}

func TestReversedPascalTriangleEdgeCases(t *testing.T) {
	if got := ReversedPascalTriangle(0); got != nil {
		t.Errorf("ReversedPascalTriangle(0) = %v, want nil", got)
	}
	if got := ReversedPascalTriangle(-1); got != nil {
		t.Errorf("ReversedPascalTriangle(-1) = %v, want nil", got)
	}

	results := ReversedPascalTriangle(1)
	if len(results) != 1 || results[0].String() != "1" || results[0].Sum != 0 {
		t.Errorf("ReversedPascalTriangle(1) = %v, want single row [1] sum 0", results)
	}
}
