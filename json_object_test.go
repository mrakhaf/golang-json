package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Firstname  string
	Lastname   string
	Middlename string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Firstname:  "Muhammad",
		Middlename: "Rakha",
		Lastname:   "Firdaus",
		Age:        22,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
