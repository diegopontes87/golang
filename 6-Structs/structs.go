package main

import "fmt"

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Struct Files")

	var newUser user
	newUser.name = "Diego"
	newUser.age = 36
	fmt.Println(newUser)

	adress := address{"Test Street", 123}
	newUser2 := user{"Renata", 36, adress}
	fmt.Println(newUser2)

	newUser3 := user{name: "Jolie"}
	fmt.Println(newUser3)

}
