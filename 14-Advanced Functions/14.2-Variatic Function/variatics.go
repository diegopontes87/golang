package main

import "fmt"

func sum(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	total := sum(1, 23, 4, 5, 7, 8, 9)
	fmt.Println(total)
	write("Hello World", 1, 2, 3, 5, 6)
}
