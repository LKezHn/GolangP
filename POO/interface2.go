package main

import "fmt"

// Structs
type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string
	loan float32
}

type Employee struct{
	Human
	company string
	money float32
}

// Inteerfaces
type Men interface{
	SayHi()
	Sing(lyrics string)
	//Guzzle(beerStein string)
} 
/*
 type YoungChap interface{
 	SayHi()
 	Sing( song string)
 	BorrowMoney(amount float32)
 }

 type ElderlyGent interface{
 	SayHi()
 	Sing(song string)
 	SpendSalary(amount float32)
 }

 func (s *Student) BorrowMoney(amount float32) {
	 s.loan += amount
 }
 
 func (e *Employee) SpendSalary(amount float32) {
	 e.money -= amount
 }
*/
func (h Human) SayHi() {
	fmt.Printf("Hi, I'm %s and my phone is %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("Singing %s\n", lyrics)
}

/* func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle, guzzle, guzzle...", beerStein)
}
*/
// Override SayHIU() for Employee
func (e Employee) SayHi() {
	fmt.Println("I'm a employee named ", e.name)
}

func main() {
	mike := Student{Human{"Mike", 25, "99534231"}, "MIT", 0.00}  
	paul := Student{Human{"Paul", 27, "89425276"}, "Harvard", 0.00} 
	sam := Employee{Human{"Sam", 36, "99342123"}, "Golang Inc.", 1000} 
	tom := Employee{Human{"Tom", 36, "88432154"}, "Things Ltd.", 5000}
	
	var i Men

	i = mike
	fmt.Println("Is Mike")
	
	i.SayHi()
	i.Sing("In the End")
	
	i = tom
	fmt.Println("Is Mike")
	i.SayHi()
	i.Sing("Paparazzi")

	mens := make([]Men, 3)

	mens[0], mens[1], mens[2] = paul, sam, mike

	for _, val := range mens{
		val.SayHi()
	}

}


