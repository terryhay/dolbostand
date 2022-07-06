package arrays_and_slices

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestArray(t *testing.T) {
	fmt.Printf("Демонстрация отличий массива от слайса\n\n")

	array := [10]int{}
	fmt.Printf("Вот массив из 10 элементов: %v\n", array)
	fmt.Printf("А вот его размер: %d\n", unsafe.Sizeof(array))
	fmt.Printf("Кстати, 10 интов имеют размер: %d\n", 10*unsafe.Sizeof(0))

	fmt.Printf("\nА как насчет функции len применительно к массиву?\nОна всякий раз будет высчитывать размер, получается?\nНе понятно, где размер хранится. Го проверять!")

	lenValue := 0
	startTime := time.Now()
	lenValue = len(array)
	workTime := time.Since(startTime)

	fmt.Printf("\n\nДля массива из %d элементов len отработала за: %v ns\n", lenValue, workTime.Nanoseconds())

	{
		bigArray := [1000]int{}
		startTime = time.Now()
		lenValue = len(bigArray)
		workTime = time.Since(startTime)

		fmt.Printf("Для массива из %d элементов len отработала за: %v ns\n", lenValue, workTime.Nanoseconds())
	}
	fmt.Printf("\nВсе-таки, размер где-то хранится. Не похоже, чтобы размер размер массива как-то жестко замедлял работу функции\nНа момент компиляции размер известен, и компилятор сохраняет его в Секретное Место,\nа при выполнении программы len ходит за этим размером")

	fmt.Println("\n\nТеперь сделаем срез и проверим адреса элементов массива и среза, хо-хо")

	fmt.Println("\nАдреса элементов массива с 1 (не 0) по 4")
	for i := 1; i < 4; i++ {
		fmt.Printf("%d: %p\n", i, &array[i])
	}

	view := array[1:4]

	fmt.Println("\nАдреса элементов среза")
	for i := 0; i < len(view); i++ {
		fmt.Printf("%d: %p\n", i, &view[i])

	}

	fmt.Printf("\nСрез хранит указатель на первый элемент массива и индексы с какого по какой\nНо маскируется под динамический массив\n")
	fmt.Printf("\nСерьезно, у среза даже capacity есть: %v\nА если есть capacity, то... в срез можно еще элементов напихать?\n", cap(view))

	fmt.Printf("%p: адрес элемента массива, который в нашем срезе сидит по 0 индексу\n\nСледим, что будет происходить с адресом нулевого элемента среза\nпри добавлении элементов в срез\n", &array[1])

	currentCap := cap(view)
	for i := 0; i < currentCap; i++ {
		view = append(view, i)
		fmt.Printf("%p: всего добавлено: %d\tlen=%d\tcap=%d\n", &view[0], i+1, len(view), cap(view))
	}

	fmt.Printf("\nМинутку! Ведь во всех этих массивах/срезах элементы лежат друг за дружкой\nАдрес нулевого элемента среза некоторое время не менялся, но ведь мы добавляли туда элементы!\nПолучается, часть элементов среза продолжает лежать в массиве, и он на них ссылается,\nа где лежат добавленные? Добавляя элементы, мы неявно редактировали исходных массив, что ли?\n\nПроверяем\n")

	view = array[1:4]
	for i := 0; i < 4; i++ {
		view = append(view, i)
	}

	for i := 0; i < len(view); i++ {
		fmt.Printf("view[%d]=%d\t%p\n", i, view[i], &view[i])
	}

	fmt.Printf("\nИсходный массив был заполнен нулями. Смотрим\n")

	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d]=%d\t%p\n", i, array[i], &array[i])
	}

	fmt.Printf("\nТак вот откуда взялась цифра capacity=9:\nэто длина массива, из которого мы сделали срез (10) минус первый индекс среза (1)\n\n")
}

func TestDynamicArray(t *testing.T) {
	dynamicArray := make([]int, 10)
	fmt.Println(dynamicArray)
	fmt.Println(unsafe.Sizeof(dynamicArray))
	fmt.Println(cap(dynamicArray))
}

func TestDynamicArrayDemo(t *testing.T) {
	fmt.Println("Dynamic array address demonstration")
	dynamicArray := make([]int, 0)

	for i := 0; i < 100; i++ {
		fmt.Printf("pointer: %p\tcap: %d\tcontent: %v\n", &dynamicArray, cap(dynamicArray), dynamicArray)
		dynamicArray = append(dynamicArray, i)
	}
}

func TestDynamicArrayFirstElementAddressDemo(t *testing.T) {
	fmt.Println("Dynamic array reallocation demonstration")
	dynamicArray := make([]int, 0)

	for i := 0; i < 100; i++ {
		dynamicArray = append(dynamicArray, i)
		fmt.Printf("pointer: %p\tcap: %d\t[0]element pointer: %p\n", &dynamicArray, cap(dynamicArray), &dynamicArray[0])
	}
}

func TestClearSlice(t *testing.T) {
	fmt.Printf("Демонстрация очистки среза\n\n")

	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = i
	}

	slice1 := slice
	fmt.Printf("%v, cap=%d\n", slice1, cap(slice1))

	fmt.Printf("Пример очистки под ноль\n")
	slice1 = nil
	fmt.Printf("%v, cap=%d\n", slice1, cap(slice1))

	fmt.Printf("Кстати, сохранились ли значения первого слайса?\n")
	fmt.Printf("%v, cap=%d\n", slice, cap(slice))

	fmt.Printf("Куда более интересный вариант с сохранением выделенного буфера\n")

	slice1 = slice
	slice1 = slice1[:0]
	fmt.Printf("%v, cap=%d\n", slice1, cap(slice1))

	fmt.Printf("А что там с исходным слайсом?\n")
	fmt.Printf("%v, cap=%d\n", slice, cap(slice))
}
