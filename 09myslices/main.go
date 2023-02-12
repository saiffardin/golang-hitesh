package main

import (
	"fmt"
)

// remove value from slice
func main() {
	fmt.Println("Welcome to class on slices")

	var courses = []string{"react","javascript","swift","python","ruby","GOLANG"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
