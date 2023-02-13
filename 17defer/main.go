package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("\t i:", i)
	}
	defer fmt.Println("\t defer in func")
	fmt.Println("\t after defer")

}

// 'defer' is function scope
