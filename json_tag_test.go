package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"img_url"`
}

func TestJSONTagArray(t *testing.T) {
	product := []Product{
		{
			Id:       "782378",
			Name:     "Rakha",
			Price:    5000,
			ImageURL: "/resource/rakha.png",
		},
		{
			Id:       "782379",
			Name:     "Firdaus",
			Price:    6000,
			ImageURL: "/resource/firdaus.png",
		},
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJSONTagArray(t *testing.T) {
	jsonRequest := `[{"id":"782378","name":"Rakha","price":5000,"img_url":"/resource/rakha.png"},{"id":"782379","name":"Firdaus","price":6000,"img_url":"/resource/firdaus.png"}]`
	jsonBytes := []byte(jsonRequest)

	products := &[]Product{}

	err := json.Unmarshal(jsonBytes, products)
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
	for _, product := range *products {
		fmt.Println(product.Id)
	}

}
