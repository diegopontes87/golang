package main

import "fmt"

func main() {
	defer function1()
	function2()
}

func function1() {
	fmt.Println("Exectuing function 1")
}

func function2() {
	fmt.Println("Exectuing function 2")
}
