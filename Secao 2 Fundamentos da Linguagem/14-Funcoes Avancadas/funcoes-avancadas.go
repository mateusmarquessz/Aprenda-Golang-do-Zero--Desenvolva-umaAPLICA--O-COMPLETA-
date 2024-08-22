package main

import "fmt"

//Retorno Nomeado
//Aqui voce ja cria as variaveis de inicio
func calculosMatematicos(n1, n2 int) (somas int, subtracao int) {
	somas = n1 + n2
	subtracao = n1 + n2
	return
}

//Funcoes Variaticas
//ela vai receber de 1 a n int numeros
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

//Funcoes Recursivas
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//Defer
func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

//Panic e Recover

func recuperarExecacao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecacao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!")
}

//Func Closure
//Basicamente sao funcoes que referenciam variaveis fora do seu corpo
func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

//funcao com ponteiros
func inverterSinal(numero int) int {
	return numero * -1
}

func inverSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

//Func Init
//Serve para executar antes da funcao main
func init() {
	fmt.Println("Executando funcao init")
}

func main() {
	somas, subtracao := calculosMatematicos(10, 20)
	fmt.Println(somas, subtracao)

	totalSoma := soma(1, 2, 3, 0)
	fmt.Println(totalSoma)

	//Funcoes Anonimas
	//Uma funcao anonima, ela funciona assim e com () eu a declaro e uso
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido ->", texto, 10)
	}("Passando Parametro")

	fmt.Println(retorno)

	//Funcao Recursiva
	posicao := uint(10)

	for i := uint(0); 1 < posicao; i++ {
		fmt.Println(fibonacci(posicao))
	}

	defer funcao1()
	//Defer = Adiar
	funcao2()

	//Panic e recover
	fmt.Println(alunoEstaAprovado(3, 6))
	fmt.Println("Pos Execucao")

	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

	//funcoes com ponteiro
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//Chamando a funcao com o ponteiro
	novoNumero := 40
	fmt.Println(novoNumero)
	inverSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
