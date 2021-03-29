package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.141592

const ( // Se pueden declarar multiples constantes a la vvez
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {

	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)


	/*
		El generador iota es usado para declarar una secuencia de valores para incrementar el valor nunmerico sin la necesidad de 
		expecificar el typo de cada uno de ellos.

		const (
			Zero = 0
			One = 1
			Two = 2
			Three = 3
			Four = 4
		)

		Uasando iota sería de la siguiente forma:
	*/
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(Two)
	fmt.Println(Four)

}