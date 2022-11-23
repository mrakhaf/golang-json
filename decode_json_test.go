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

func TestDecodeJSONArray(t *testing.T) {
	jsonRequest := `{"Firstname":"Muhammad","Lastname":"Firdaus","Middlename":"Rakha","Age":22,"Married":false,"Hobbies":["Football","Music","Automotive"]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.Lastname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)

	for _, hobby := range customer.Hobbies {
		fmt.Println(hobby)
	}
}
