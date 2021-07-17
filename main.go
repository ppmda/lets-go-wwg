package main

import "fmt"

func main() {
	hora := 35

	switch {
	case hora >= 6 && hora < 12:
		fmt.Printf("Agora são %d horas e é de manhã", hora)

	case hora >= 12 && hora < 18:
		fmt.Printf("Agora são %d horas e é de tarde", hora)

	case hora >= 18 && hora < 24:
		fmt.Printf("Agora são %d horas e é de noite", hora)

	case hora >= 0 && hora < 6:
		fmt.Printf("Agora são %d horas e é de madruga", hora)
	default:
		fmt.Printf("O valor %d informado não é uma hora válida", hora)
	}
}
