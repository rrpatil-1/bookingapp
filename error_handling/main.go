package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator should not be zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("started error handling in function")
	ans, _ := divide(10, 0)

	fmt.Println("division of number is:", ans)

}
