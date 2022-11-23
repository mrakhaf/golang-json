package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, err := os.Open("Customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	customer := Customer{}
	err = decoder.Decode(&customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
}
