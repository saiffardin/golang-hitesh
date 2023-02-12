package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	hitesh := User{"Saif", "saiffardin@gmail.com", true, 27}
	fmt.Println(hitesh)
	fmt.Println(hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
