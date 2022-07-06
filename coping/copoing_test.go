package coping

import (
	"fmt"
	"testing"
)

func TestCoping(t *testing.T) {
	fmt.Printf("Демонстрация особенностей копирования\n\n")

	fmt.Printf("Обычные переменные копируются целиком\n")

	{
		v := 1
		copyV := 2
		copyV = v

		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &v, v, &copyV, copyV)
		fmt.Println("Изменим значение copyV")
		copyV++

		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &v, v, &copyV, copyV)
	}

	{
		fmt.Printf("\nПосмотрим, как поведет себя статический массив\n")

		arr := [3]int{1, 2, 3}
		copyArr := [3]int{0, 0, 0}

		arr = copyArr
		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &arr, arr, &copyArr, copyArr)

		fmt.Println("Изменим значение copyArr")
		copyArr[0] = 4
		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &arr, arr, &copyArr, copyArr)
	}

	{
		fmt.Printf("\nПосмотрим, как поведет себя срез\n")

		slice := []int{1, 2, 3}
		copySlice := []int{0, 0, 0}

		copySlice = slice
		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &slice, slice, &copySlice, copySlice)

		fmt.Println("Изменим значение copySlice")
		copySlice[0] = 4
		fmt.Printf("v[%p]=%d, copyV[%p]=%d\n", &slice, slice, &copySlice, copySlice)

		fmt.Println("Скопировалась direct-part, вместе со ссылкой на буфер. Он общий у двух срезов")
	}
}
