package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays in go have fixed sizes
	var array1 [5]string
	array1[0] = "first"
	fmt.Println(array1)

	array2 := [5]string{"position1", "position2", "position3", "position4", "position5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices are dynamic structures
	slice := []int{10, 11, 9, 8}
	fmt.Println(slice)

	// Slices and Arrays are different types
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	// Adding new values to a Slice
	slice = append(slice, 202)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Slices has the same behavior as a pointer
	array2[1] = "Changed Position"
	fmt.Println(slice2)

	// Internal Arrays
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
