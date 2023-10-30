package main

import "fmt"

func main() {
	//explicit declaration
	var variable1 string = "Variable 1"
	//implicit declaration
	variable2 := "Variable 2"
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "Variable 3"
		variable4 string = "Variable 4"
	)

	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variable 5", "Variable 6"
	fmt.Println(variable5, variable6)

	//const values
	const constant1 string = "Contant 1"
	fmt.Println(constant1)

	//inverting variables values
	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
