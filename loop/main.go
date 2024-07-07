package main

import "fmt"

func main() {
	for i := 0; i <= 1; i++ {
		fmt.Println(i)
	}

	counter := 0

	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 1 {

			break
		}
	}

	number := []int{4, 5, 6, 7}

	for index, value := range number {
		fmt.Printf("index of number is: %v, and value is %d\n", index, value)
	}

	data := "hello, world!"
	fmt.Println("\n")

	for index, char := range data {
		fmt.Printf("index of of data is %d and value is %c\n", index, char)

	}
}
