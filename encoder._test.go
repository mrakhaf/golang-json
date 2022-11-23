package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, err := os.Create("sample_output.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Rakha",
		Lastname:   "Firdaus",
	}
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
