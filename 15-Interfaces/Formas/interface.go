package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func area(f forma) {
	fmt.Println("A área é de: ", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type circulo struct {
	raio float64
}

func main() {

	r := retangulo{10, 10}
	c := circulo{5}

	area(r)
	area(c)

}
