package main

import (
	"fmt"
)

func main() {
	i := 0
	fmt.Println("Incrementing i")
	for i < 10 {
		i++
		fmt.Println(i)
	}

	fmt.Println("Incrementing j")
	for j := 0; j < 10; j += 1 {
		fmt.Println(j)
	}

	names := []string{"Diego", "Renata", "Douglas"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name: ":      "Diego",
		"last name: ": "Borges",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}
