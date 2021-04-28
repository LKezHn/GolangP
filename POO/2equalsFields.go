package main

import "fmt"

type Human struct{
	name string
	age int
	weight int
	phone string // Campo phone en Human
}

type Employee struct{
	Human
	speciality string
	phone string // Campo phone en Student
}

func main(){

	Bob := Employee{ Human{"Bob", 32, 150, "8975-8799"}, "Chef", "9923-5423"}

	/*
		En caso de dos campos con el mismo nombre Go toma al campo que tenga el nivel mas alto, en este caso el campo phone de Employee.
	*/
	fmt.Printf("Bob's Employee phone number: %v\n", Bob.phone) 
	fmt.Printf("Bob's Human phone number: %v\n", Bob.Human.phone)

}