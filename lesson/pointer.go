package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func updateUserByValue(u User) {
	u.Name = "Updated Name"
	fmt.Println("Inside updateUserByValue:", u)
}

func updateUserByValueAndReturn(u User) User {
	u.Name = "Updated Name"
	fmt.Println("Inside updateUserByValue:", u)

	return u
}

func updateUserByPointer(u *User) {
	u.Name = "Updated Name"
	fmt.Println("Inside updateUserByPointer:", *u)
}

func main() {
	user := User{Name: "Alice", Age: 25}

	fmt.Println("Before updateUserByValue:", user)
	updateUserByValue(user)                       // Passing by value (copy)
	fmt.Println("After updateUserByValue:", user) // Original remains unchanged

	fmt.Println("\nBefore updateUserByValueAndReturn:", user)
	updateUserByValue(user)                       // Passing by value (copy)
	fmt.Println("After updateUserByValueAndReturn:", user) // We see the change

	fmt.Println("\nBefore updateUserByPointer:", user)
	updateUserByPointer(&user)                      // Passing by pointer
	fmt.Println("After updateUserByPointer:", user) // Original is changed
}
