package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
