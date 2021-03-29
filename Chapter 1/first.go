package main

import "fmt"

var z = "cadena con formato"

func main(){
	y := "James Bond" // := Solo se usa para definir la variable por primera vez
	x := 7

	fmt.Println(x, y)
	fmt.Printf("Esta es una %v.\n", z)  // %v = valor de variable
	fmt.Printf("El tipo de dato de z es %T.\n", z) // %T = tipo de dato de variable

	x = 10 // = Se usa para reasignar el valor 

	fmt.Println(x)
	hello()
}

func hello(){
	fmt.Println(z)
}