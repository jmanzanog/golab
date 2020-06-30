package _interface

type Animal interface {
	Respirar()
}

func EjecutarRespirar(animal Animal) {
	animal.Respirar()
}
