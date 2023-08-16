package main

import (
	"fmt"
)

func main() {
	// Arithmetics
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 9 * 2
	divisionRest := 10 % 2

	fmt.Println(sum, subtraction, division, multiplication, divisionRest)

	// Atributions
	var variable1 string = "String"
	variable2 := "String2"
	fmt.Println(variable1, variable2)

	// Relationals
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Logicals
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!false)

	// Unary
	number := 10
	number++
	number += 15 // number = number + 15
	fmt.Println(number)

	number--
	number -= 20 // number = number - 20
	fmt.Println(number)

	number *= 3 // number = number * 3
	number /= 3 // number = number * 3

	// Ternary
	var text string
	if number > 5 {
		text = "Bigger than 5"
	} else {
		text = "Smaller than 5"
	}
	fmt.Println(text)
}
