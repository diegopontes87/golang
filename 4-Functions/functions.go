package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculations(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	var sum int8 = sum(10, 20)
	fmt.Println(sum)

	var f = func(text string) {
		fmt.Println(text)
	}
	f("Function f")

	var sum2, subtraction2 int8 = calculations(10, 15)
	fmt.Println(sum2, subtraction2)

	sum3, _ := calculations(10, 15)
	fmt.Println(sum3)
}
