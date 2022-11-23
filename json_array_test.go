package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONArray(t *testing.T) {
	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Rakha",
		Lastname:   "Firdaus",
		Age:        22,
		Married:    false,
		Hobbies:    []string{"Football", "Music", "Automotive"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

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

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Rakha",
		Lastname:   "Firdaus",
		Age:        22,
		Married:    false,
		Hobbies:    []string{"Football", "Music", "Automotive"},
		Address: []Address{
			{
				Street:     "Brawijaya V",
				Country:    "Indonesia",
				PostalCode: "15810",
			},
			{
				Street:     "Brawijaya VI",
				Country:    "Indonesia",
				PostalCode: "15810",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJSONArrayComplex(t *testing.T) {
	jsonRequest := `{"Firstname":"Muhammad","Lastname":"Firdaus","Middlename":"Rakha","Age":22,"Married":false,"Hobbies":["Football","Music","Automotive"],"Address":[{"Street":"Brawijaya V","Country":"Indonesia","PostalCode":"15810"},{"Street":"Brawijaya VI","Country":"Indonesia","PostalCode":"15810"}]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer.Address)

	for _, address := range customer.Address {
		fmt.Println(address.Country)
	}

}
func TestStringOnlyDecodeJSONArrayComplex(t *testing.T) {
	jsonRequest := `[{"Street":"Brawijaya V","Country":"Indonesia","PostalCode":"15810"},{"Street":"Brawijaya VI","Country":"Indonesia","PostalCode":"15810"}]`
	jsonBytes := []byte(jsonRequest)

	address := &[]Address{}

	err := json.Unmarshal(jsonBytes, address)

	if err != nil {
		panic(err)
	}

	fmt.Println(address)

	for _, adr := range *address {
		fmt.Println(adr.Country)
	}
}
