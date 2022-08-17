package main
import "fmt"

func main() {
	i := 10

	name, middleName := "akshay", "rao"

	fmt.Println((name + " " + middleName))

	newName, newLastName := nameFunc()
	fmt.Println(newName + " " + newLastName)

	WhatWasSaid, otherthing := something()
	fmt.Println(WhatWasSaid, otherthing)

	defer fmt.Println(i) // execute this line at the end
	fmt.Println("Hello World!")
}

func nameFunc() (string, string) {
	a := "brodi"
	b := "yoo"
	return a, b
}

func something() (string, int) {
	return "Omkar", 10
}
