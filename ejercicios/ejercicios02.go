package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


var num int
var err error

func Mult() (string) {
	scanner := bufio.NewScanner(os.Stdin)
	
	var texto string
	for {
		println("Ingrese un numero: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {break}
		}
	}

	for i:=1; i<=10; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", num, i, i*num)
	}

	return texto
}