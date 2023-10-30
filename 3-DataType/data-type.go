package main

import (
	"errors"
	"fmt"
)

func main() {
	//INT
	//int8, int 16, int32, int64
	var number int = 12310
	fmt.Println(number)

	//int without signal
	var number2 uint = 12310
	fmt.Println(number2)

	//alias
	//int32 = rune
	var number3 rune = 12310
	fmt.Println(number3)

	//int8 = byte
	var number4 byte = 123
	fmt.Println(number4)

	//FLOATS - REAL NUMBERS
	var realNumber1 float32 = 123.34
	fmt.Println(realNumber1)
	var realNumber2 float32 = 1214132434.34123123123
	fmt.Println(realNumber2)

	//STRINGS
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text 2"
	fmt.Println(str2)

	// there isn't chat in golang. It brings the int from ASC table
	char := 'B'
	fmt.Println(char)

	//BOOLEAN
	var boolean1 bool = true
	var boolean2 bool = false
	fmt.Println(boolean1)
	fmt.Println(boolean2)

	//ERROR
	var err error = errors.New("Internal error")
	fmt.Println(err)
}
