package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered")
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverExecution()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("The media is exactly 6")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("Post exectuion")
}
