package pointers

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
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

func getFakeDataArray() []Data {
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

func GetDataAndWritePointer() Data {
	data := NewFakeData()
	fmt.Printf("%p, %p: создали объект с таким указателем (на всякий случай, и адрес первого поля покажем)\n", &data, &data.ID)
	return data
}

func TestPointers(t *testing.T) {
	fmt.Println("Меняется ли указатель при передаче по значению?")
	data := GetDataAndWritePointer()
	fmt.Printf("%p, %p: ну, вот какими они стали\n", &data, &data.ID)

}
