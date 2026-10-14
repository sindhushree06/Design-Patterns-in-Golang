package main

import "fmt"

type student struct {
	name string
}

func (s *student) clone() (student student) {
	return *s
}

func main() {
	originalObj := student{name: "yash"}

	clonedObj := originalObj.clone()
	clonedObj.name = "sindhu"
	fmt.Println(originalObj)
	fmt.Println(clonedObj)
}
