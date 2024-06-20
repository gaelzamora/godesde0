package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("ingrese numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		}
	}

	fmt.Println("ingrese numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		}
	}

	fmt.Println("ingrese Leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}