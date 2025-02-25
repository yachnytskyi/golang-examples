package main

import "fmt"

func modifySlice(s []int) {
	fmt.Println("Inside function:", s) // Output: [0, 0, 0]
	s[0] = 100                         // Modify an element
}

func addElement(s []int) {
	s = append(s, 4)
	fmt.Println("Inside function:", s) // Output: [1, 2, 3, 4]
}

func modifyAndReturnSlice(s []int) []int {
	fmt.Println("Inside function:", s) // Output: [0, 0, 0]
	s[0] = 100
	s = append(s, 4)
	return s
}

func addElementByPointer(s *[]int) {
	fmt.Println("Inside function:", *s) // Output: [1, 2, 3]
	*s = append(*s, 4)
}

func main() {
	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println("Outside function:", mySlice) // Output: [100, 2, 3]

	mySlice = make([]int, 3, 4)
	addElement(mySlice)
	fmt.Println("Outside function:", mySlice) // Output: [1, 2, 3]

	mySlice = make([]int, 3, 4)
	mySlice = modifyAndReturnSlice(mySlice)
	fmt.Println("Outside function:", mySlice) // Output: [100, 0, 0, 4]

	newSlice := []int{1, 2, 3}
	addElementByPointer(&newSlice)
	fmt.Println("Outside function:", newSlice) // Output: [1, 2, 3, 4]
}
