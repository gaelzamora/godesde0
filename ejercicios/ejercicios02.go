package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func Mult() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("ingrese el numero a multiplicar: ")
	if scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto "+err.Error())
		} else {
			for i:=1; i <= 10; i++ {
				fmt.Println(num, " * ", i, " = ", num*i)
			}
		}
	}
}