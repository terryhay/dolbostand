package maps

import (
	"fmt"
	"testing"
)

func TestMaps(t *testing.T) {
	fmt.Printf("Демонстрация особенностей мэпов\n")

	map1 := make(map[int]int)
	for i := 0; i < 3; i++ {
		map1[i] = i
	}

	map2 := map1

	fmt.Printf("Вот два мэпа\n")
	fmt.Printf("%p: %v\n", &map1, map1)
	fmt.Printf("%p: %v\n", &map2, map2)

	fmt.Printf("\nОбщий ли у них буфер? Добавим десяток значений, например\n")

	for i := 3; i < 13; i++ {
		map2[i] = i
	}
	fmt.Printf("%p: %v\n", &map1, map1)
	fmt.Printf("%p: %v\n", &map2, map2)

	fmt.Printf("Получается, общий (копируются мэпы только через цикл)\n")
}
