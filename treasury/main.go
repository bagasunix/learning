package main

import "fmt"

func main() {
	fmt.Println("=== Question 1 ===")
	fmt.Println(Question1A([]int{-1, 1}))      // 0
	fmt.Println(Question1A([]int{-1, -7, -5})) // -2
	fmt.Println(Question1A([]int{1, 2, 1, 6})) // 3

	fmt.Println(Question1B([]int{-1, 1}))      // 0
	fmt.Println(Question1B([]int{-1, -7, -5})) // -1
	fmt.Println(Question1B([]int{1, 2, 1, 6})) // 1

	fmt.Println("=== Question 2 ===")
	r, err := Question2([]int{0, 1, 0}, 1000, 1010)
	fmt.Println(r, err) // 0 not found
	r, err = Question2([]int{0, 1, 0}, 1259, 2601)
	fmt.Println(r, err) // 62952 nil

	fmt.Println("=== Question 3 ===")
	fmt.Println(Question3([]int{1, 2}, []int{1, 3}))             // [1]
	fmt.Println(Question3([]int{1, 2, 2}, []int{1, 2, 4}))       // [2 1]
	fmt.Println(Question3([]int{1, 2, 2, 4, 5}, []int{1, 2, 4})) // [4, 2, 1]

	fmt.Println("=== Question 4 ===")
	fmt.Println(Question4(1223334)) // map[1:1 2:2 3:3 4:1]

	fmt.Println("=== Question 5 ===")
	fmt.Println(Question5(1, 2, 3, 4))    // 14
	fmt.Println(Question5(1, 2, 3, 4, 5)) // 39
}
