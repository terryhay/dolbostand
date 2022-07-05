package main

import (
	"testing"
)

var contractPtr *Contract
var contractObj Contract

func BenchmarkMustGetContractAsPointer(b *testing.B) {
	var contract *Contract
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contract = mustGetContractAsPointer()
	}
	contractPtr = contract
}

func BenchmarkMustGetContractAsObject(b *testing.B) {
	var contract Contract
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contract = mustGetContractAsObject()
	}
	contractObj = contract
}

func BenchmarkMustInitContractPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contract := Contract{}
		mustInitContractAsPointer(&contract)
		contractObj = contract
	}
}

func BenchmarkMustInitContractObject(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contract := new(Contract)
		mustInitContractAsPointer(contract)
		contractPtr = contract
	}
}
