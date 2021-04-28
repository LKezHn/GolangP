package main

import (
	"fmt"
)

type Skills []string

type Human struct{
	name string
	age int
	weight int
}

type Student struct{
	Human // Struct embebida
	Skills // Tipo Slice embebido
	int // Tipo de dato entero anónimo
	speciality string
}

func main(){

	jane := Student{ Human: Human{"Jane",21,110}, speciality: "Biology"}

	fmt.Printf("Her name is: %v\n", jane.name)
	fmt.Printf("Her age is: %v\n", jane.age)
	fmt.Printf("Her weight is: %v\n", jane.weight)
	fmt.Printf("Her speciality is: %v\n", jane.speciality)
	
	// Modificando el valor de las skills
	jane.Skills = []string{"Anatomy"}
	fmt.Printf("Her skills are: %v\n", jane.Skills)
	fmt.Println("Append two new skills")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Printf("Her new skills are: %v\n", jane.Skills)
	
	// Modificando el valor int anonimo
	jane.int = 24
	fmt.Printf("Her favorite number is: %v\n", jane.int)
}