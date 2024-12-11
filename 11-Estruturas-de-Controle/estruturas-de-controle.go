package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDasemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
		//Executa a prox linha
		fallthrough
	case num == 2:
		return "Segunda"

	default:
		return "Num Invalid"
	}

}

func main() {

	num := 10

	if num > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	//Variavel usando If init, so presente dentro do If.
	if num2 := num; num2 > 0 {
		fmt.Println("É maior que 0")
	}

	//	SWITCH
	dia := diaDaSemana(3)

	fmt.Println(dia)

	dia2 := diaDaSemana(10)

	fmt.Println(dia2)

}
