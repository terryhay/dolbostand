// Демонстрация сравнения времени работы циклов
package loops_speed_comparison

import "testing"

var arr = func(n uint) (arr []int) {
	arr = make([]int, 0, n)
	for i := 0; i < int(n); i++ {
		arr = append(arr, i)
	}
	return arr
}(1000)

func StandardLoop(arr []int) int {
	res := 0
	for _, element := range arr {
		res += element
	}
	return res
}

func IndexLoop(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res
}

func EmptyStatementLoop(arr []int) int {
	var (
		res   int
		i     int
		count = len(arr)
	)
	for {
		res += arr[i]
		if i++; i == count {
			break
		}
	}
	return res
}

func CPPLoop(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func BenchmarkStandardLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardLoop(arr)
	}
}

func BenchmarkIndexLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IndexLoop(arr)
	}
}

func BenchmarkEmptyStatementLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EmptyStatementLoop(arr)
	}
}

func BenchmarkCPPLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CPPLoop(arr)
	}
}
