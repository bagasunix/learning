package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func innerSum(row []int) int {
	sum := 0
	for i := 1; i < len(row)-1; i++ {
		sum += row[i]
	}
	return sum
}

func formatRow(row []int) string {
	parts := make([]string, len(row))
	for j, v := range row {
		parts[j] = strconv.Itoa(v)
	}
	return strings.Join(parts, " ")
}

func parity(v int) string {
	if v%2 == 0 {
		return "gn"
	}
	return "gj"
}

// center menempatkan s di tengah kolom selebar width.
func center(s string, width int) string {
	pad := width - len(s)
	left := pad / 2
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", pad-left)
}

func RunCustomPascal(n int) {
	if n <= 0 {
		return
	}

	tri := make([][]int, n)
	for i := 0; i < n; i++ {
		tri[i] = make([]int, i+1)
		tri[i][0], tri[i][i] = 1, 1
		for j := 1; j < i; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}

	w := len(strconv.Itoa(tri[n-1][(n-1)/2])) + 3

	for i := n - 1; i >= 0; i-- {
		line := strings.Repeat(" ", (n-1-i)*w/2)
		sum := 0

		for j, v := range tri[i] {
			token := strconv.Itoa(v)
			if j > 0 && j < i { // bagian dalam
				token += parity(v)
				sum += v
			}
			line += center(token, w)
		}

		fmt.Printf("%-*s== %d (Baris %d)\n", n*w, line, sum, i+1)
	}
}

func RunCustomPascalSimple(n int) {
	if n <= 0 {
		return
	}

	tri := make([][]int, n)
	for i := range n {
		tri[i] = make([]int, i+1)
		tri[i][0], tri[i][i] = 1, 1
		for j := 1; j < i; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}

	for i := n - 1; i >= 0; i-- {
		sum := 0
		for j, v := range tri[i] {
			if j == 0 || j == i {
				fmt.Printf("%d ", v)
				continue
			}
			sum += v
			fmt.Printf("%d%s ", v, parity(v))
		}
		fmt.Printf("== %d (Baris %d)\n", sum, i+1)
	}
}

func main() {
	RunCustomPascal(6)
	fmt.Println(" ")
	RunCustomPascalSimple(6)
	fmt.Println(" ")

	// Contoh penggunaan triangle Pascal dengan jumlah baris yang berbeda
	pascal := generate(15)
	for _, a := range slices.Backward(pascal) {
		for _, b := range a {
			fmt.Print(b)
			fmt.Print(" ")
		}
		fmt.Printf("== %d\n", innerSum(a))
	}
	fmt.Println(" ")
	for _, a := range pascal {
		for _, b := range a {
			fmt.Print(b)
			fmt.Print(" ")
		}
		fmt.Printf("== %d\n", innerSum(a))
	}
}
