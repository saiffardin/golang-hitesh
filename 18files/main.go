package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "This needs to go in a file"
	file, err := os.Create("./devops.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)


	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./devops.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)


	fmt.Println(string(databyte))
}

func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}
