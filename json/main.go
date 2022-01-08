package main

import (
	"encoding/json"
	"log"
)

type person struct {
	Firstname string `json:"first_name"` //from json
	Lastname  string `json:"last_name"`
	Superhero bool   `json:"superhero"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"superhero": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"superhero": true
		}	
	]`

	var unmarshalled []person //because functions used for json are called marshalled and unmarshalled

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error Unmarshalling:", err)
	}

	log.Println(unmarshalled)

	// Write Json from a Struct

	var myslice []person

	m1 := person{
		Firstname: "Diana",
		Lastname:  "Prince",
		Superhero: false,
	}

	myslice = append(myslice, m1)

	m2 := person{
		Firstname: "Barry",
		Lastname:  "Allen",
		Superhero: true,
	}

	myslice = append(myslice, m2)

	json2, err := json.Marshal(myslice)
	// To Indent use marshalIndent
	json3, err := json.MarshalIndent(myslice, "", "")

	log.Println(string(json2))
	log.Println(string(json3))
}
