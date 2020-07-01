package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//formatos imprimir
	variable := "valor"

	fmt.Printf("tipo %T\n", 3)
	fmt.Printf("valor %v\n", 3)
	fmt.Printf("valor base 10 %d\n", 500)
	fmt.Printf("valor base hex %x\n", "hexa")
	fmt.Printf("valor float rep %f\n", 12.303030)
	fmt.Printf("valor cientifica rep %E\n", 3.14159626667267)
	fmt.Printf("valor puntero rep %p\n", &variable)
	fmt.Printf("especios fijo nuemeros |%6d|%6d|\n", 12, 345)
	fmt.Printf("numero de cifras |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("justificar salida |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("numero de letras |%6s|%6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	fmt.Println(uno(5))
	fmt.Println(dos(5))
	_, valor1 := dos(10)
	fmt.Println(valor1)
	fmt.Println(tres(1, 2, 3, 4, 5, 6))

	tabla := multiplicar(9)

	for i := 0; i < 10; i++ {
		fmt.Println(tabla())
	}
	fmt.Println(tabla())

	array1 := []int{2, 22, 2, 2, 2, 2, 2, 2, 2, 34, 5, 5, 8, 2, 2, 3, 4, 5, 6, 6}
	array2 := array1[5:6]
	array3 := make([]int, 1, 5)
	array3 = append(array3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Printf("array3 largo %d , capacidad %d\n", len(array3), cap(array3))
	fmt.Printf("array2 largo %d , capacidad %d\n", len(array2), cap(array2))
	fmt.Printf("array1 largo %d , capacidad %d\n", len(array1), cap(array1))

	manipularMapas()
	goto RUTINA
RUTINA:
	fmt.Println("rutina")
}

func uno(number int) (z int) {
	z = number + 2
	return z
}
func dos(number int) (E bool, value int) {
	value = number * 2
	return true, value
}
func tres(numbers ...int) (s string) {
	for _, value := range numbers {
		s += strconv.Itoa(value)
	}

	return s
}

func multiplicar(valor int) func() int {
	number := valor
	sequence := 0

	return func() int {
		sequence++
		return sequence * number
	}
}

func manipularMapas() {
	numbersPrime := make(map[int]bool, 10)
	numbersPrime[2] = true
	numbersPrime[15] = false
	fmt.Println(numbersPrime)
	paisesCapitales := map[string]string{

		"Colombia":  "Bogota",
		"Argentina": "Buenos Aires",
		"España":    "Madrid",
		"France":    "Paris"}
	delete(paisesCapitales, "Colombia")
	fmt.Println(paisesCapitales)
	for pais, capital := range paisesCapitales {
		fmt.Print(capital, " es capital de ", pais, "\n")
	}

}
