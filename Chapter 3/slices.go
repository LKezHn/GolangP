package main

import "fmt"

func main(){

	aSlice := []int{1, 2, 3, 4} // Un slice se crea igual que un array pero sin especificar el numero de elementos que éste contiene.
	fmt.Println(aSlice[1:3]) // La notacion [:] excluye al indice que va despues de ":"

	s1 := make([]int, 5) // Creaando un array usando la funcion make(), make() solo funciona en slices, maps y channels.
	reSlice := s1[1:3] 
	fmt.Println(s1) 
	fmt.Println(reSlice) 
	
	reSlice[0] = -100 
	reSlice[1] = 123456 

	// Multidimensional Slice
	s2 := make([][]int, 3)

	fmt.Println(s1) 
	fmt.Println(reSlice)
}