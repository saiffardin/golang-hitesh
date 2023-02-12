package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to class on pointers")
	// var ptr *int
	// fmt.Println("value of pointer is:", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("memory address:", ptr)
	fmt.Println("value:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("NEW - memory address:", ptr)
	fmt.Println("NEW - value:", *ptr)

}
