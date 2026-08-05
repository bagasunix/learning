# Learning - Live Coding Interview Practice (Golang)

Kumpulan soal latihan live coding interview dalam Go.

## Struktur

    Learning/
    ├── go.mod
    │
    ├── [ Level: Mudah ]
    ├── fizzbuzz/
    ├── fibonacci/
    ├── palindrome/
    ├── anagram/
    ├── reverse_string/
    ├── factorial/
    ├── count_vowels/
    │
    ├── [ Level: Sedang ]
    ├── two_sum/
    ├── valid_parentheses/
    ├── prime_number/
    ├── binary_search/
    ├── merge_sorted/
    ├── max_subarray/
    ├── first_unique_char/
    ├── linked_list_cycle/
    ├── stack_queue/
    ├── climb_stairs/
    ├── majority_element/
    ├── rotate_array/
    ├── string_compression/
    │
    ├── solutions/
    │   └── solutions.go         # Jawaban referensi lengkap
    └── README.md

## Cara Pakai

### 1. Kerjakan soal

Buka file .go di folder soal (misal `fizzbuzz/fizzbuzz.go`), baca deskripsi,
lalu implementasikan fungsi yang ada `// TODO`.

### 2. Jalankan test untuk soal tertentu

    cd ~/Learning
    go test ./fizzbuzz/ -v

### 3. Jalankan semua test sekaligus

    cd ~/Learning
    go test ./... -v

### 4. Lihat jawaban referensi

Buka `solutions/solutions.go` untuk implementasi lengkap semua soal.

## Daftar Soal

### Level: Mudah

| No | Soal             | Konsep                          |
|----|------------------|---------------------------------|
| 1  | FizzBuzz         | Loop, kondisional, modulo       |
| 2  | Fibonacci        | Loop, recursion, memoization    |
| 3  | Palindrome       | String manipulation, two-pointer|
| 4  | Anagram          | Sorting / hash map              |
| 5  | Reverse String   | Slicing, recursion              |
| 6  | Factorial        | Loop, recursion                 |
| 7  | Count Vowels     | String, set/map                 |

### Level: Sedang

| No | Soal               | Konsep                              |
|----|--------------------|-------------------------------------|
| 8  | Two Sum            | Nested loop, hashmap O(n)           |
| 9  | Valid Parentheses  | Stack                               |
| 10 | Prime Number       | Math, Sieve of Eratosthenes         |
| 11 | Binary Search      | Divide & conquer, recursion         |
| 12 | Merge Sorted       | Two-pointer, merge step             |
| 13 | Max Subarray       | Kadane's algorithm (DP)             |
| 14 | First Unique Char  | Hash map, two-pass                  |
| 15 | Linked List Cycle  | Floyd's tortoise & hare (O(1) space)|
| 16 | Stack Queue        | Stack-based queue, amortized O(1)   |
| 17 | Climb Stairs       | DP / fibonacci variant              |
| 18 | Majority Element   | Boyer-Moore voting (O(1) space)     |
| 19 | Rotate Array       | 3-reverse trick (O(1) space)        |
| 20 | String Compression | Run-length encoding                |

## Setup Go

Jika Go belum terinstall, jalankan:

    # Download dan extract ke home directory
    cd /tmp
    wget https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
    mkdir -p ~/go-sdk && tar -C ~/go-sdk -xzf go1.24.6.linux-amd64.tar.gz

    # Tambahkan ke PATH (tambahkan ke ~/.bashrc untuk permanen)
    export PATH=$HOME/go-sdk/go/bin:$HOME/go/bin:$PATH

## Tips Live Coding Interview

1. Pahami soal dulu - Jangan buru-buru coding, tanya klarifikasi
2. Pikirkan edge cases - Input kosong, n=0, n=1, negative
3. Bicara sambil coding - Jelaskan pikiranmu ke interviewer
4. Mulai dari brute force - Lalu optimasi kalau diminta
5. Test jawabanmu - Jalankan test sebelum bilang selesai
6. Perhatikan complexity - Tahu Big-O dari solusimu
