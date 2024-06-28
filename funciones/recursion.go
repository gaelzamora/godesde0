package funciones

import "fmt"

func Exponencia(value int) {
	if value > 100000000 {
		return
	}
	fmt.Println(value)
	Exponencia(value*2)
} 