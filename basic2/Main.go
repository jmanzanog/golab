package main

import (
	_interface "./interface"
	_struct "./struct"
	"fmt"
)

type PersonaExt struct {
	_struct.Person
}

func main() {
	p := PersonaExt{_struct.Person{Name: "Jose", Age: 30}}

	p1 := new(_struct.Person)
	p1.Name = "Carlos"
	p2 := new(PersonaExt)
	p2.NewPerson("Jonas", 17)
	p2.NewPerson("Adam", 50)
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)

	poches := new(_interface.Gato)
	poches.Nombre = "Poches"
	_interface.Animal.Respirar(poches)

	jose := _interface.Humano{Nombre: "Jose"}

	_interface.Animal.Respirar(&jose)
	_interface.Stuff.Respirar(&jose)

}
