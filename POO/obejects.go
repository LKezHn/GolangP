package main

import (
	"fmt"
)
// Con structs simulamos las clases
type Person struct{
	Name string
	LastName string
	Age int
}

// Simulando métodos de una clase
func (p *Person) getFullName() string {
	return p.Name + " " + p.LastName
}

func main(){
	unknow := Person{
		"Luis", 
		"Martinez", 
		21, // Siempre agregar la coma al final
	}

	// fmt.Println(unknow.Name)
	// fmt.Println(unknow.LastName)
	// fmt.Println(unknow.Age)
	fmt.Println(unknow.getFullName())
}