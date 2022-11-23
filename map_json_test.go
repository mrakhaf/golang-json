package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMapJSONDecode(t *testing.T) {
	jsonRequest := `{"id":"12456","name":"Apple Iphone", "price":12000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(reflect.TypeOf(result["price"]).Kind())
}

func TestMapJSONEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Apple Macbook Pro",
		"price": 600000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
