package main

import "fmt"

func main() {
	fmt.Print("learnig array in go")
	// var name [5]string
	// name[0] = "prince"
	// name[1] = "akash"
	// name[2] = "ravi"
	// name[4] = "rahul"

	// fmt.Println("\n Name of person is:", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is:", numbers)

	fmt.Println("length of array is:", len(numbers))
	var price [5]int

	fmt.Println(price)

	var name [5]string

	name[2] = "rahul"

	name[0] = "prince"

	fmt.Printf("Name of array %q", name)

}
