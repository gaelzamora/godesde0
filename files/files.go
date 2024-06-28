package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gaelzamora/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var text string = ejercicios.Mult()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err.Error())
		return
	}
	fmt.Fprintln(archivo, text)
	archivo.Close()
}

func SumaTabla() {
	var text string = ejercicios.Mult()
	if !Append(filename, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, text string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append "+err.Error())
		return false
	}

	_, err = arch.WriteString(text)
	
	if  err != nil {
		fmt.Println("Error durante el WriteString "+err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo "+err.Error())
		return
	}

	fmt.Println(string(archivo))
}