package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//Serve para separar as funcoes que serao testas
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area Recebida %f e diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area Recebida %f e diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

}
