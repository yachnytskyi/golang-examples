package main

import (
	"fmt"
	"testing"
)

func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

func BenchmarkModifyAndReturnSlice(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := generateSlice(size)
				modifyAndReturnSlice(s)
			}
		})
	}
}

func BenchmarkAddElementByPointer(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := generateSlice(size)
				addElementByPointer(&s)
			}
		})
	}
}
