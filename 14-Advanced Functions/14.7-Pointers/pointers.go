package main

import "fmt"

func signalInverter(number int) int {
	return number * -1
}

func signalInverterPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	invertedNumber := signalInverter(number)
	fmt.Println(invertedNumber)
	fmt.Println(number)

	newNumber := 40
	signalInverterPointer(&newNumber)
	fmt.Println(newNumber)
}
