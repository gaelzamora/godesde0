package main

import (
	"fmt"

	"github.com/gaelzamora/godesde0/ejercicios"
)

func main() {
	value, cadena := ejercicios.FuncionRegresa("10")

	fmt.Println(value)
	fmt.Println(cadena)
}