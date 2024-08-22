package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	//Todos os dados do meu array tem que ser do mesmo tipo

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays internos
	//make ela cria um array de 10 mas so me mostra as 8 posicoes
	//isso e um array interno
	//quando o slice estourar a capacidade o go vai dobrar o array dele
	slice3 := make([]int, 8, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
