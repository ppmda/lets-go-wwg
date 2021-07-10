package main

import "fmt"

func main() {
	cores := map[string]string{
		"cinza": "#424242",
		"roxo":  "#a378f8",
	}
	fmt.Println("1", cores)

	delete(cores, "cinza")
	fmt.Println("2", cores)

	//azul, existe := cores["azul"]
	//fmt.Println(azul)
	//fmt.Println(existe)
}
