package main

import "fmt"

func mathCalculations(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, subtraction := mathCalculations(1, 2)
	fmt.Println(sum, subtraction)
}
