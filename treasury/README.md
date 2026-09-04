# Treasury — Golang Coding Test

Kumpulan 5 soal coding test Golang beserta solusi dan penjelasannya.

## Cara Jalankan

```bash
# Run semua test
go test ./... -v

# Jalankan main (lihat output semua soal)
go run .
```

---

## Question 1A — Missing Number (Hash Set)

```go
func Question1A(arr []int) int
```

### Aturan

| Kondisi | Return |
|---|---|
| Array mengandung bilangan positif DAN negatif | `0` |
| Array hanya berisi bilangan positif | Bilangan bulat positif terkecil yang **tidak ada** di array |
| Array hanya berisi bilangan negatif | Bilangan bulat negatif terbesar (paling dekat ke 0) yang **tidak ada** di array |

### Contoh

```
[-1, 1]        → 0    (ada positif dan negatif → return 0)
[-1, -7, -5]   → -2   (hanya negatif, ada: {-1,-5,-7}, yang hilang paling dekat 0: -2)
[1, 2, 1, 6]   → 3    (hanya positif, ada: {1,2,6}, yang hilang terkecil: 3)
[5]            → 1    (hanya positif, 1 belum ada)
[-3]           → -1   (hanya negatif, -1 belum ada)
[1, 2, 3]      → 4    (semua dari 1 ada, lanjut ke 4)
[-1, -2, -3]   → -4   (semua dari -1 ada, lanjut ke -4)
```

### Cara Kerja

Gunakan hash set untuk lookup O(1), lalu scan linear mulai dari 1 atau -1:

**All positive** — cari bilangan yang "hilang" mulai dari 1:
```
[1, 2, 1, 6] → present = {1, 2, 6}
cek 1 → ada
cek 2 → ada
cek 3 → TIDAK ADA → return 3
```

**All negative** — cari bilangan yang "hilang" mulai dari -1:
```
[-1, -7, -5] → present = {-1, -5, -7}
cek -1 → ada
cek -2 → TIDAK ADA → return -2
```

---

## Question 1B — Extreme Value (Min/Max)

```go
func Question1B(arr []int) int
```

### Aturan

Varian berbeda dari soal yang sama — alih-alih mencari angka yang **hilang**, cari angka yang **ada** di array paling dekat ke 0:

| Kondisi | Return |
|---|---|
| Array mengandung bilangan positif DAN negatif | `0` |
| Array hanya berisi bilangan positif | Nilai **terkecil** (minimum) di array |
| Array hanya berisi bilangan negatif | Nilai **terbesar** (maksimum, paling dekat ke 0) di array |

### Contoh

```
[-1, 1]        → 0    (mixed → return 0)
[-1, -7, -5]   → -1   (hanya negatif, max = -1)
[1, 2, 1, 6]   → 1    (hanya positif, min = 1)
[5]            → 5
[-3]           → -3
[-7, -5, -3]   → -3   (paling dekat ke 0)
```

### Cara Kerja

Single pass tanpa extra memory — track min atau max sesuai kondisi:

```
[-1, -7, -5]
hasNeg = true → cari max
  result = -1
  cek -7 → -7 < -1, skip
  cek -5 → -5 < -1, skip
return -1
```

---

## Question 2 — Reverse & Exclude

```go
func Question2(exclude []int, nums ...int) (int, error)
```

### Aturan

1. Gabungkan semua angka di `nums` menjadi satu string
2. Balik (reverse) seluruh digit
3. Hapus digit yang ada di `exclude`
4. Kalau hasilnya kosong → return `0, error("not found")`

### Contoh

```
Question2([0,1,0], 1000, 1010) → 0, "not found"
Question2([0,1,0], 1259, 2601) → 62952, nil
Question2([],      12,   34)   → 4321, nil   (exclude kosong → tidak ada yang dihapus)
```

### Cara Kerja

**Contoh 1:**
```
Join:    "1000" + "1010" = "10001010"
Reverse: "01010001"
Exclude: {0, 1} → hapus semua 0 dan 1
Sisa:    "" → return 0, "not found"
```

**Contoh 2:**
```
Join:    "1259" + "2601" = "12592601"
Reverse: "10629521"
Exclude: {0, 1} → hapus digit 0 dan 1
         1→hapus, 0→hapus, 6→simpan, 2→simpan, 9→simpan, 5→simpan, 2→simpan, 1→hapus
Sisa:    "62952" → return 62952, nil
```

**Contoh 3 (exclude kosong):**
```
Join:    "12" + "34" = "1234"
Reverse: "4321"
Exclude: {} → tidak ada yang dihapus
Sisa:    "4321" → return 4321, nil
```

---

## Question 3 — Intersection Descending

```go
func Question3(numsA, numsB []int) []int
```

### Aturan

1. Cari angka di `numsA` yang **juga ada** di `numsB`
2. Duplikat diabaikan (tiap angka hanya muncul sekali di hasil)
3. Urutkan **descending** (besar ke kecil)

### Contoh

```
numsA=[1,2],     numsB=[1,3]     → [1]
numsA=[1,2,2],   numsB=[1,2,4]   → [2, 1]
numsA=[5,6],     numsB=[1,2]     → []      (tidak ada irisan)
numsA=[3,1,2],   numsB=[1,2,3]   → [3,2,1]
```

### Cara Kerja

```
numsA=[1,2,2], numsB=[1,2,4]

Buat set dari numsB: {1, 2, 4}

Iterasi numsA:
  1 → ada di set? YA → tambah ke result
  2 → ada di set? YA → tambah ke result
  2 → sudah pernah masuk? YA → skip (dedup)

result = [1, 2]
Sort descending → [2, 1]
```

---

## Question 4 — Digit Frequency

```go
func Question4(num int) map[int]int
```

### Aturan

1. Pecah angka menjadi digit-digit tunggal
2. Hitung berapa kali setiap digit muncul
3. Return sebagai `map[digit]count`
4. Angka negatif diperlakukan seperti positif (tanda minus diabaikan)

### Contoh

```
1223334 → map[1:1 2:2 3:3 4:1]
```

### Cara Kerja

```
1223334 → digits: [1, 2, 2, 3, 3, 3, 4]

Hitung:
  1 → 1x
  2 → 2x
  3 → 3x
  4 → 1x

Result: map[1:1 2:2 3:3 4:1]
```

---

## Question 5 — Concurrent Multiply & Sum

```go
func Question5(nums ...int) int
```

### Aturan

1. Pasangkan elemen: index 0 dengan 1, index 2 dengan 3, dst
2. Kalau elemen terakhir tidak punya pasangan → kalikan dengan dirinya sendiri
3. Setiap perkalian dijalankan **secara concurrent** (goroutine)
4. Penjumlahan hasil perkalian dilakukan di **main thread** (setelah semua goroutine selesai)

### Contoh

```
[1, 2, 3, 4]    → (1×2) + (3×4)         = 2 + 12      = 14   (2 goroutine)
[1, 2, 3, 4, 5] → (1×2) + (3×4) + (5×5) = 2 + 12 + 25 = 39   (3 goroutine)
[2, 3]          → (2×3)                  = 6                   (1 goroutine)
[7]             → (7×7)                  = 49                  (self-multiply)
[]              → 0                                            (kosong)
```

### Cara Kerja

```
nums = [1, 2, 3, 4, 5]

Buat pasangan:
  pair[0] = (1, 2)
  pair[1] = (3, 4)
  pair[2] = (5, 5)  ← tidak punya pasangan, self-multiply

Jalankan 3 goroutine secara concurrent:
  goroutine 0: results[0] = 1 × 2 = 2
  goroutine 1: results[1] = 3 × 4 = 12
  goroutine 2: results[2] = 5 × 5 = 25

Main thread tunggu semua selesai (wg.Wait())
Main thread jumlahkan: 2 + 12 + 25 = 39
```

### Kenapa Concurrency?

Perkalian tiap pasangan **independen** satu sama lain — tidak saling bergantung.
Goroutine memungkinkan semua perkalian berjalan paralel, baru hasilnya dikumpulkan di main thread.

---

## Struktur File

```
treasury/
├── questions.go       semua implementasi (Question1A, Question1B, Question2-5)
├── questions_test.go  semua unit test
├── main.go            entry point, contoh pemanggilan
└── README.md          penjelasan ini
```
