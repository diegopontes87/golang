package main

import "fmt"

func main() {
	result := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("parameter")

	fmt.Println(result)
}
