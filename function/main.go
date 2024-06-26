package main

import "fmt"

func SimpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function")
	SimpleFunction()
	ans := add(3, 2)
	data := multiply(4, 3)
	fmt.Println("sum of 2 number is:", ans)
	fmt.Println("multiplication of number is:", data)
}
