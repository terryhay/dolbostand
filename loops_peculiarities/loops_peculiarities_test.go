// Демонстрация особенностей циклов
package loops_peculiarities

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"testing"
	"time"
)

type Data struct {
	ID       string `json:"id"`
	Index    int    `json:"index"`
	GUID     string `json:"guid"`
	IsActive bool   `json:"isActive"`
	Balance  string `json:"balance"`
	Picture  string `json:"picture"`
	Age      int    `json:"age"`
	EyeColor string `json:"eyeColor"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
}

func NewFakeData() Data {
	return Data{
		ID:       gofakeit.Color(),
		Index:    int(gofakeit.Int32()),
		GUID:     gofakeit.Color(),
		IsActive: gofakeit.Bool(),
		Balance:  gofakeit.Color(),
		Picture:  gofakeit.Color(),
		Age:      int(gofakeit.Uint16()),
		EyeColor: gofakeit.Color(),
		Name:     gofakeit.Name(),
		Gender:   gofakeit.Color(),
	}
}

func getFakeDataSlice() []Data {
	return []Data{
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
		NewFakeData(),
	}
}

func NewFakeDataPointer() *Data {
	return &Data{
		ID:       gofakeit.Color(),
		Index:    int(gofakeit.Int32()),
		GUID:     gofakeit.Color(),
		IsActive: gofakeit.Bool(),
		Balance:  gofakeit.Color(),
		Picture:  gofakeit.Color(),
		Age:      int(gofakeit.Uint16()),
		EyeColor: gofakeit.Color(),
		Name:     gofakeit.Name(),
		Gender:   gofakeit.Color(),
	}
}

func getFakeDataPointerArray() []*Data {
	return []*Data{
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
		NewFakeDataPointer(),
	}
}

func TestPointers(t *testing.T) {
	fmt.Printf("Сейчас пробежимся по массиву объектов и указателей, посмотрим, есть ли что занятного\n\n")

	fmt.Println("Во-первых, поменяется ли адрес переменной obj, когда мы обходим слайс с объктами?")
	objectSlice := getFakeDataSlice()
	for _, obj := range objectSlice {
		fmt.Printf("%p: %s\n", &obj, obj.ID)
	}

	fmt.Printf("\nАдрес не меняется: элементы слайса копируются.\n\nЕсли у нас в одном случае 10 структур с 10 полями, а в другом 10 указателей на структуры с 10 полями...\nВо сколько раз дольше будет работать обход слайса, в котором лежат объеты?\n")

	startTime := time.Now()
	for _, obj := range objectSlice {
		obj.ID = ""
	}
	objLoopWorkTime := time.Since(startTime)

	pointerSlice := getFakeDataPointerArray()

	startTime = time.Now()
	for _, pointer := range pointerSlice {
		pointer.ID = ""
	}
	pointerLoopWorkTime := time.Since(startTime)

	fmt.Printf("\nВ %.2f раз дольше\n", float64(objLoopWorkTime.Nanoseconds())/float64(pointerLoopWorkTime.Nanoseconds()))
}

func TestGorutine(t *testing.T) {
	fmt.Printf("При обходе слайса объектов, они копируются в один объект.\nЭто приведет к забавным последствиями, если мы будем передавать этот объект в горутину\n\n")

	wg := sync.WaitGroup{}

	fmt.Printf("Запустим цикл с захватом объекта горутинами и выводом поля ID\n")
	objectSlice := getFakeDataSlice()
	for _, obj := range objectSlice {
		wg.Add(1)
		go func() {
			fmt.Printf("%p: %s\n", &obj, obj.ID)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("\nКстати, а какой это по счету элемент массива? Вот последний\n[9]: %s\n", objectSlice[9].ID)
	fmt.Printf("\nЦикл отработал быстро и наплодил 10 горутин. Все они захватили один и тот же объект.\nЕсли повезет, парочка горутин успеет напечатать значение, еще во время работы цикла.\nНо подавляющее их большинство протупит цикл, а когда запустится, то захваченное ими значение будет равно\nпоследнему значению слайса\n")
}
