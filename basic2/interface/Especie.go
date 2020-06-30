package _interface

import "fmt"

type Humano struct {
	Nombre string
}

func (h *Humano) Respirar() {
	fmt.Print(h.Nombre, " Respirando\n")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Respirar() {
	fmt.Print(g.Nombre, " Respirando\n")
}
