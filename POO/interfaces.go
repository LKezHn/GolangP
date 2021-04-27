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
	Person
	Role string
	Languages []string
}

type Student interface{
	hello() string
}

func (p Person) hello() string {
	return "Hi! I'm "+ p.Name
}

func main(){

	p1 := Developer{
		Person{ "Luis", "Martinez", 22},
		"FullStack Developer",
		[]string{"Golang", "Python", "Javasccript",},
	}

	p2 := Person{"Joe", "Collins", 32}
	p3 := Person{"Lord", "Inestroza", 33}

	students := []Student{p1, p2, p3}

	for _, v := range students {
		fmt.Println(v.hello())
	}
}