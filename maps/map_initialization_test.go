package maps

import "testing"

const mapSize = 1000

// go: noinline
func initMapWithoutReserve(size int) map[int]int {
	res := make(map[int]int)
	for i := 0; i < size; i++ {
		res[i] = i
	}

	return res
}

// go: noinline
func initMapWithReserve(size int) map[int]int {
	res := make(map[int]int, size)
	for i := 0; i < size; i++ {
		res[i] = i
	}

	return res
}

func BenchmarkInitMapWithoutReserve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		initMapWithoutReserve(mapSize)
	}
}

func BenchmarkInitMapWithReserve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		initMapWithReserve(mapSize)
	}
}
