package main

import (
	"io"
	"os"
)

func main(){
	myString := ""
	arguments := os.Args // Args para capturar argumentos enviados desde la consola
	
	if len(arguments) == 1{
		myString = "I need one argument"
	}else{
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString) // os.Stdout = salida estandar del sistema.
	io.WriteString(os.Stdout, "\n") // io.WriteString = escribir el texto de myString en un io.Writer pasado como primer parametro.
}