package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 10}
		areaEsperada := float64(100)

		areaRecebida := ret.area()

		if areaRecebida != areaEsperada {
			t.Errorf("Area incorreta, %f", areaRecebida)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.area()

		if areaEsperada != areaRecebida {
			t.Errorf("Area incorreta, %f", areaRecebida)
		}

	})

}
