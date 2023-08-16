package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	person1 := person{"Diego", "Borges", 36, 184}
	fmt.Println(person1)

	student1 := student{person1, "Design", "PUC"}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.height)
}
