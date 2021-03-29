package main

import (
	"bufio" // Usado de forma comun para inputs y outputs de archivos.
	"fmt"
	"os"
)

func main(){

	var f *os.File // Creamos un archivo
	f = os.Stdin // Obtenenmos el input desde consola
	defer f.Close() // Instruccion se ejecuta al final de la funcion

	scanner := bufio.NewScanner(f) // Creamos un scanner usadno bufio.NewScanner y pasamos el archivo
	for scanner.Scan(){
		fmt.Println(">", scanner.Text()) // Por cada escaneo que haga el scanner se imprime el texto leido.
	}
}