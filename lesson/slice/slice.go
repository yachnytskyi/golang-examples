package main

import "fmt"

func modifyArray(arr [3]int) {
    arr[0] = 99
    fmt.Println("Inside function:", arr) 
}

func modifySlice(s []int) {
	fmt.Println("Inside function:", s) 
	s[0] = 100                        
}

func addElement(s []int) {
	s = append(s, 4)
	s = append(s, 5)
	fmt.Println("Inside function:", s) 
}

func modifyAndReturnSlice(s []int) []int {
	fmt.Println("Inside function:", s) 
	s[0] = 100
	s = append(s, 4)
	return s
}

func addElementByPointer(s *[]int) {
	fmt.Println("Inside function:", *s)
	*s = append(*s, 4)
}

func main() {
	myArray := [3]int{1, 2, 3}
    modifyArray(myArray)
    fmt.Println("Outside function:", myArray) 

	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println("Outside function:", mySlice) 

	mySlice = make([]int, 3, 4)
	addElement(mySlice)
	fmt.Println("Outside function:", mySlice)

	mySlice = make([]int, 3, 4)
	mySlice = modifyAndReturnSlice(mySlice)
	fmt.Println("Outside function:", mySlice) 

	newSlice := []int{1, 2, 3}
	addElementByPointer(&newSlice)
	fmt.Println("Outside function:", newSlice) 
}
