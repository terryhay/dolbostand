package main

import (
	"encoding/json"
	"fmt"
)

const contractJSON = `{
	"_id": "5d2f4fcf76c35513af00d47e",
	"index": 1,
	"guid": "ed687a14-590b-4d81-b0cb-ddaa857874ee",
	"isActive": true,
	"balance": "$3,837.19",
	"picture": "http://placehold.it/32x32",
	"age": 28,
	"eyeColor": "green",
	"name": "Rochelle Espinoza",
	"gender": "female",
	"company": "PARLEYNET",
	"email": "rochelleespinoza@parleynet.com",
	"phone": "+1 (969) 445-3766",
	"address": "956 Little Street, Jugtown, District Of Columbia, 6396",
	"about": "Excepteur exercitation labore ut cupidatat laboris mollit ad qui minim aliquip nostrud anim adipisicing est. Nisi sunt duis occaecat aliquip est irure Lorem irure nulla tempor sit sunt. Eiusmod laboris ex est velit minim ut cillum sunt laborum labore ad sunt.\r\n",
	"registered": "2016-03-20T12:07:25 -00:00",
	"latitude": 61.471517,
	"longitude": 54.01596,
	"greeting": "Hello, Rochelle Espinoza!You have 9 unread messages.",
	"favoriteFruit": "banana"
}`

type Contract struct {
	ID            string  `json:"_id"`
	Index         int     `json:"index"`
	GUID          string  `json:"guid"`
	IsActive      bool    `json:"isActive"`
	Balance       string  `json:"balance"`
	Picture       string  `json:"picture"`
	Age           int     `json:"age"`
	EyeColor      string  `json:"eyeColor"`
	Name          string  `json:"name"`
	Gender        string  `json:"gender"`
	Company       string  `json:"company"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	Address       string  `json:"address"`
	About         string  `json:"about"`
	Registered    string  `json:"registered"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Greeting      string  `json:"greeting"`
	FavoriteFruit string  `json:"favoriteFruit"`
}

//go:noinline
func mustGetContractAsPointer() *Contract {
	contract := new(Contract)
	if err := json.Unmarshal([]byte(contractJSON), contract); err != nil {
		panic(err)
	}

	return contract
}

//go:noinline
func mustGetContractAsObject() Contract {
	contract := Contract{}
	if err := json.Unmarshal([]byte(contractJSON), &contract); err != nil {
		panic(err)
	}

	return contract
}

//go:noinline
func mustInitContractAsPointer(contract *Contract) {
	if err := json.Unmarshal([]byte(contractJSON), contract); err != nil {
		panic(err)
	}
}

func main() {
	contract := mustGetContractAsObject()
	fmt.Println(contract)
	contractPtr := mustGetContractAsPointer
	fmt.Println(contractPtr())

	newContractObj := Contract{}
	newContractObjPtr := &newContractObj
	fmt.Printf("%p", newContractObjPtr)
}
