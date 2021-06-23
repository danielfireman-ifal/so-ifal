package main

import (
	"fmt"
	"time"
)

type item struct {
	nome     string
	telefone string
}

// variável global (área de memória compartilhada)
var lista = []item{
	{"Marcos Soares", "8212345678"},
	{"Thiago Gazaroli", "8256781234"},
}

// A função escritora adiciona um item a lista.
func escritora() {
	lista = append(lista, item{"Ismael Feliz", "8298765678"})
}

func leitora() {
	fmt.Printf("%v\n", lista)
}

func main() {
	fmt.Printf("%v\n", lista)

	go escritora()
	go leitora()

	// vamos esperar que as outras threads executem.
	// isso é uma forma de coordenação. A thread principal
	// espera pelas threads criadas.
	time.Sleep(10 * time.Millisecond)

	// Ao executar várias vezes, vai ficar claro a necessidade
	// de coordenação, principalmente quando a comunicação
	// acontece via memória compartilhada.
	fmt.Printf("%v\n", lista)
}
