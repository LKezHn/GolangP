package main

import (
	"fmt"
)

type Person struct{
	Name string
	LastName string
	Age int
}

type Developer struct{
	Person // Funciona como herencia
	Role string
	Languages []string
}

func (p *Person) hi() string {
	return "Hello! I'm " + p.Name
}

func main(){

	none := Developer{
		Person{"Luis", "Martinez", 22},
		"FullStack Developer",
		[]string{"Golang", "Python", "Javascript",},
	} 


	fmt.Println(none.Name)
	fmt.Println(none.Languages)
	fmt.Println(none.Role)
	fmt.Println(none.hi())
}