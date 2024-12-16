package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Nome  string `json:"nome"`
	Idade uint8  `json:"idade"`
}

func main() {
	usr := user{"Marcos", 21}

	userJSON, err := json.Marshal(usr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userJSON)
	fmt.Println(bytes.NewBuffer(userJSON))

}
