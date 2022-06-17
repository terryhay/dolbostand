package arrays_and_slices

import (
	"testing"
)

const sliceSize = 100

func makeSlice1(n int) []int {
	slice := []int{}
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}

func makeSlice2(n int) []int {
	slice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}

func makeSlice3(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}

func BenchmarkSliceInitialization1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSlice1(sliceSize)
	}
}

func BenchmarkSliceInitialization2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSlice2(sliceSize)
	}
}

func BenchmarkSliceInitialization3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSlice3(sliceSize)
	}
}
