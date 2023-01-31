package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string   `json:"Name"`
	Age  int      `json:"Age"`
	Tags []string `json:"Tags"`
}

func main() {

	DecodeJsonFirstWay()

}

func GetReq() {}

func PostReq() {}

func EncodeJson() {
	person := []person{
		{"mahmoud", 23, []string{"Dev", "BackEnd"}},
	}
	finalJson, err := json.MarshalIndent(person, "", "\t")
	if err != nil {

		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJsonFirstWay() {

	jsonData := []byte(`
        {
                "name": "mahmoud",
                "age": 23,
                "tags": ["Dev","BackEnd"]
		}
		`)
	var person person

	isValid := json.Valid(jsonData)
	if isValid {
		fmt.Println("Data is Valid")
		err := json.Unmarshal(jsonData, &person)
		if err != nil {
			return
		}
		fmt.Printf("%#v\n", person)
	} else {
		fmt.Println("DATA IS NOT VALID")
	}

}

func DecodeJsonSecondWay() {

	jsonData := []byte(`
        {
                "name": "mahmoud",
                "age": 23,
                "tags": ["Dev","BackEnd"]
		}
		`)

	var person map[string]interface{}
	isValid := json.Valid(jsonData)
	if isValid {
		fmt.Println("Data is Valid")
		err := json.Unmarshal(jsonData, &person)
		if err != nil {
			return
		}
		fmt.Printf("%#v\n\n", person)
		for k, v := range person {
			fmt.Printf("Key is %v and Value is %v", k, v)

		}

	} else {
		fmt.Println("DATA IS NOT VALID")
	}

}