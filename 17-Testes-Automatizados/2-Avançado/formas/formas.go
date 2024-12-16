package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func Area(f Forma) {
	fmt.Println("A área é de: ", f.area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

type Circulo struct {
	Raio float64
}
