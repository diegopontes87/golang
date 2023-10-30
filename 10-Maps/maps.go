package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Diego",
		"lastName": "Pontes",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Diego",
			"last":  "Pontes",
		},
		"address": {
			"street": "Amparo",
			"number": "123",
		},
	}
	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

}
