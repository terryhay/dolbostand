package comparing

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type Contract struct {
	Title  string
	Button *Button
}

type Button struct {
	Title      string
	ActionName string
}

func makeContract() *Contract {
	return &Contract{
		Title: "Красная кнопка",
		Button: &Button{
			Title:      "Не нажимай меня, пожалуйста",
			ActionName: "crashSisAction",
		},
	}
}

func TestContract(t *testing.T) {
	fmt.Printf("Демонстрация сравнения\nСравниваются самые нижние объекты иерархии побайтово\n\n")

	testContract := &Contract{
		Title: "Красная кнопка",
		Button: &Button{
			Title:      "Не нажимай меня, пожалуйста",
			ActionName: "crashSisAction",
		},
	}
	contract := makeContract()

	fmt.Println(testContract == contract)
	fmt.Printf("%p != %p, значение указателей разное и неважно, что по ним лежат одинаковые по сути значения\n\n",
		testContract,
		contract)

	testContractObj := *testContract
	contractObj := *contract

	fmt.Println(testContractObj == contractObj)
	fmt.Printf("%p != %p, ну и что? Ну разыменовали указатели, и теперь у нас сравниваются два объекта,\nу которых разные значения указателей на кнопки\n\n",
		testContractObj.Button,
		contractObj.Button)

	fmt.Println(cmp.Equal(testContract, contract))
	fmt.Printf("А теперь заработало. Остается вопрос, действительно ли функция делает полный обход?\nПоменяем что-нибудь в кнопке\n")

	testContract.Button.Title = "Ладно, нажимай"
	fmt.Println(cmp.Equal(testContract, contract))

	fmt.Printf("Да, теперь false\n")
}
