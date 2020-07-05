package main

import "fmt"

func main() {
	numero := 5
	var puntero *int
	puntero = &numero
	numero = 6
	puntero2 := &puntero
	fmt.Printf("valor de numero: %v\n", numero)
	fmt.Printf("direccion de numero: %v\n", &numero)
	fmt.Printf("valor puntero de direccion de numero: %v\n", *puntero)
	fmt.Printf("valor puntero2 que a su vez apunta a puntero 1: %v\n", **puntero2)

	fmt.Printf("numero sin incrementar: %v\n", numero)
	incremento(numero)
	fmt.Printf("numero incrementado: %v\n", numero)

	fmt.Printf("numero sin incrementar: %v\n", numero)
	incremento2(&numero)
	fmt.Printf("numero incrementado: %v\n", numero)

}

func incremento(numero int) {
	numero++
	fmt.Printf("incremento de numero: %v\n", numero)
}

func incremento2(numero *int) {
	*numero++
	fmt.Printf("incremento de numero: %v\n", *numero)
}
