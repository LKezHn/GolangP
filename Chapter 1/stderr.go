package main

import (
	"os"
	"io"
)

func main(){
	myString := ""
	argunments := os.Args

	if len(argunments) == 1 {
		myString = "Necesito un parámetro"
	}else{
		myString = argunments[1]
	}

	io.WriteString(os.Stdout, "Todo está bien\n")
	io.WriteString(os.Stderr, myString) // Imprimiendo mensaje de error
	io.WriteString(os.Stderr, "\n")
}