package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	assistant.Write()
	erro := checkmail.ValidateFormat("test.com")
	fmt.Println(erro)
	erro2 := checkmail.ValidateFormat("test@gmail.com")
	fmt.Println(erro2)
}
