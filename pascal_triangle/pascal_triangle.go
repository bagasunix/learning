package pascaltriangle

import (
	"fmt"
	"strconv"
	"strings"
)

// =============================================================================
// REVERSED PASCAL'S TRIANGLE + GENAP/GANJIL LABEL + ROW SUM
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Pascal's triangle construction, string formatting
// Complexity: O(n²) time, O(n²) space
// =============================================================================
//
// SOAL:
//   1. Bangun segitiga Pascal n baris.
//   2. Angka di dalam baris (selain dua angka 1 di ujung kiri/kanan) diberi
//      akhiran "gn" (genap) atau "gj" (ganjil) sesuai paritasnya.
//   3. Angka-angka bukan ujung dijumlahkan per baris.
//   4. Segitiga dibalik: baris dengan elemen terbanyak di atas, baris ke-1
//      di paling bawah.
//
// CONTOH (n=4):
//   Baris 1: 1                 -> sum 0
//   Baris 2: 1 1                -> sum 0
//   Baris 3: 1 2gn 1             -> sum 2
//   Baris 4: 1 3gj 3gj 1          -> sum 6
//
//   ReversedPascalTriangle(4) mengembalikan urutan: baris4, baris3, baris2, baris1.
//
// CARA KERJA:
//   Baris berikutnya dibangun dari baris sebelumnya: row[i] = prev[i-1]+prev[i],
//   ujung selalu 1. Label ujung (index 0 dan terakhir) tetap angka biasa,
//   sisanya diberi suffix genap/ganjil. Sum hanya menjumlahkan elemen non-ujung
//   (baris dengan <=2 elemen otomatis sum=0). Setelah semua baris terbentuk,
//   urutan slice dibalik.
//
// JEBAKAN UMUM:
//   - Baris dengan 1 atau 2 elemen tidak punya elemen "dalam" -> sum harus 0,
//     bukan index out of range.
//   - Label ujung TIDAK diberi suffix genap/ganjil meski nilainya genap.
// =============================================================================

// RowResult menyimpan label tiap elemen (dengan suffix gn/gj untuk elemen dalam)
// dan jumlah elemen non-ujung pada baris tersebut.
type RowResult struct {
	Labels []string
	Sum    int
}

// String menggabungkan label satu baris dipisah spasi, contoh: "1 3gj 3gj 1".
func (r RowResult) String() string {
	return strings.Join(r.Labels, " ")
}

// buildTriangle membangun segitiga Pascal n baris (urutan normal, baris 1 dulu).
func buildTriangle(n int) [][]int {
	rows := make([][]int, n)
	var prev []int
	for i := 0; i < n; i++ {
		row := nextRow(prev)
		rows[i] = row
		prev = row
	}
	return rows
}

// nextRow menghitung baris pascal berikutnya dari baris sebelumnya.
func nextRow(prev []int) []int {
	if prev == nil {
		return []int{1}
	}
	row := make([]int, len(prev)+1)
	row[0], row[len(row)-1] = 1, 1
	for i := 1; i < len(row)-1; i++ {
		row[i] = prev[i-1] + prev[i]
	}
	return row
}

// labelRow memberi suffix "gn"/"gj" untuk elemen selain dua ujung.
func labelRow(row []int) []string {
	labels := make([]string, len(row))
	for i, v := range row {
		if i == 0 || i == len(row)-1 {
			labels[i] = strconv.Itoa(v)
			continue
		}
		suffix := "gj"
		if v%2 == 0 {
			suffix = "gn"
		}
		labels[i] = strconv.Itoa(v) + suffix
	}
	return labels
}

// rowSum menjumlahkan elemen bukan ujung pada satu baris (0 kalau baris <=2 elemen).
func rowSum(row []int) int {
	sum := 0
	for i := 1; i < len(row)-1; i++ {
		sum += row[i]
	}
	return sum
}

// ReversedPascalTriangle mengembalikan n baris segitiga Pascal (dengan label
// genap/ganjil dan sum per baris) dalam urutan terbalik: baris ke-n dulu,
// baris ke-1 terakhir.
func ReversedPascalTriangle(n int) []RowResult {
	if n <= 0 {
		return nil
	}
	rows := buildTriangle(n)
	results := make([]RowResult, n)
	for i, row := range rows {
		results[i] = RowResult{Labels: labelRow(row), Sum: rowSum(row)}
	}
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}

// PrintReversedPascalTriangle mencetak segitiga terbalik, rata tengah
// berdasarkan lebar baris terlebar (baris ke-n, yang berada di atas).
func PrintReversedPascalTriangle(n int) {
	results := ReversedPascalTriangle(n)
	if len(results) == 0 {
		return
	}
	width := len(results[0].String())
	for _, r := range results {
		line := r.String()
		pad := (width - len(line)) / 2
		fmt.Printf("%s%s  (sum=%d)\n", strings.Repeat(" ", pad), line, r.Sum)
	}
}
