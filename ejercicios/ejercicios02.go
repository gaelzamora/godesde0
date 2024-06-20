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
		fmt.Printf("%d x %d = %d\n", num, i, i*num)
	}
}