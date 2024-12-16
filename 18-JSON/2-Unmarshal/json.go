package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Nome  string `json:"nome"`
	Idade uint8  `json:"idade"`
}

func main() {

	usrJSON := `{"nome": "Marcos", "idade":21}`

	var usr user

	if err := json.Unmarshal([]byte(usrJSON), &usr); err != nil {
		log.Fatal(err)
	}

	fmt.Println(usr)
}
