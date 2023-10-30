package main

import "fmt"

func main() {
	fmt.Println("Control Structs")
	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less or equal 15")
	}

	if number2 := number; number2 > 0 {
		fmt.Println("Number is greater than 0")
	}
}
