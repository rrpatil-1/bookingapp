package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	Month := "January"

	switch Month {
	case "January", "Feb", "march":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}

	temp := -10
	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp >= 0 && temp < 10:
		fmt.Println("cold")
	case temp >= 10 && temp < 20:
		fmt.Println("Cool")
	case temp >= 20 && temp < 30:
		fmt.Println("warm")
	default:
		fmt.Println("Hot")
	}

}
