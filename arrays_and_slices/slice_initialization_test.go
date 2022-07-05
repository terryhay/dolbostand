package arrays_and_slices

import (
	"fmt"
	"testing"
)

const sliceSize = 1000

func makeSlice1(n int) []int {
	slice := make([]int, 0)
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

func TestPrintCapacity(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 10001; i++ {
		if i < 11 || (i < 101 && i%10 == 0) || (i < 1001 && i%100 == 0) || (i > 1001 && i%1000 == 0) {
			fmt.Printf("%d %d\n", i, cap(slice))
		}
		slice = append(slice, i)
	}
}
