package main

import "fmt"

func main() {
	fmt.Println("maps")

	usr := map[string]string{
		"name":      "Marcos",
		"sobrenome": "Goulart",
	}

	fmt.Println(usr)
	fmt.Println(usr["name"])

	usr2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Joao",
		},
		"curso": {
			"nome":   "CC",
			"campus": "Paulista",
		},
	}

	fmt.Println(usr2)

	delete(usr2, "curso")

	fmt.Println(usr2)

	usr2["gender"] = map[string]string{
		"type": "male",
	}

	fmt.Println(usr2)

}
