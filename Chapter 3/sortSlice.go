package main

import (
	"fmt"
	"sort"
)

type myStruct struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]myStruct, 0)
	mySlice = append(mySlice, myStruct{ "Luis", 175, 50 } ) 
	mySlice = append(mySlice, myStruct{ "Enrrique", 180, 80}) 
	mySlice = append(mySlice, myStruct{ "Matt", 170, 90}) 
	mySlice = append(mySlice, myStruct{ "Joe", 160, 85}) 
	fmt.Println(mySlice)

	/* 
		sort.Slice() solo funciona en Go 1.8 en adelante. 
		sort.Slice() recibe como primer parámetro el slice a ordenar, y como segundo parametro la funcion que contiene el criterio
		por el cual será ordenado el Slice
	*/
	sort.Slice(mySlice, func(i, j int) bool { 
		return mySlice[i].height < mySlice[j].height
	})

	fmt.Println("Ordenando de menor a mayor")
	fmt.Println(mySlice)
	
	sort.Slice(mySlice, func( i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	
	fmt.Println("Ordenando de mayor a menor")
	fmt.Println(mySlice)
}