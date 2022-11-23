package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"Firstname":"Muhammad","Lastname":"Rakha"}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Lastname)

}
