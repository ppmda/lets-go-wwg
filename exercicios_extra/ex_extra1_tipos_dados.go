package main

import "fmt"

func main() {
	// Exercicio 1

	// O erro ocorre porque esta sendo utilizado o int8, com o int ele considera os sinais negativos e prositivos.
	// Então nesse caso ele esta considerando 0 ao 127 e do -1 ao -128
	// Se alterarmos a variavel para uint8, o quilometro e impresso na tela.
	// Outra forma de imprimir, é alterando o tipo da variavel para int64
	var quilometros uint8
	quilometros = 150

	fmt.Println(quilometros)
}
