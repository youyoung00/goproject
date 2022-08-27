package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	fmt.Println("스택메모리와 힙메모리 시작!!")
	fmt.Println()

	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
