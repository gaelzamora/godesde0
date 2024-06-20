package ejercicios

import (
	"fmt"
	"strconv"
)


func FuncionRegresa(value string) (int, string) {
	valueInt, error := strconv.Atoi(value)

	if error != nil{
		fmt.Println("Error en la conversion: ", error)
	} else if valueInt > 100 {
		return valueInt, "Es mayor a 100"
	} else {
		return valueInt, "Es menor a 100"
	}

	return valueInt, ""
}