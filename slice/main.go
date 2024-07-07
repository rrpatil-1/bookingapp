package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 4, 5}
	// numbers = append(numbers, 1, 5, 6, 7, 8, 9)
	// fmt.Println(numbers)
	// fmt.Printf("data type of silce %T\n", numbers)
	// fmt.Println("leng of slice", len(numbers))

	fmt.Println("slice:", numbers)
	fmt.Println("len:", len(numbers))
	fmt.Println("capacity:", cap(numbers))

	stock := make([]int, 4)

	fmt.Println("slice:", stock)
	fmt.Println("len:", len(stock))
	fmt.Println("capacity:", cap(stock))
}
