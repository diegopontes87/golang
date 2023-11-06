package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user %s!\n", u.name)
}

func main() {
	user1 := user{"Diego", 36}
	fmt.Println(user1)
	user1.save()
}
