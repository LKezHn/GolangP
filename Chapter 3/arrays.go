package main

import (
	"fmt"
)

func main(){
	var arr [8]int // Para un array vacío se necesita usar var.
	otherArr := [3] int{9, 15, 21} // Se puede usar := para declarar arrays con al menos un valor inicial.

	douArr := [2][2] int{ // Array multidiemnsional se declara igual que en C.
		{ 1,2 },
		{ 3,4 }, // Al final de ultimo elemento siempre debe ir una coma.
	}

	count := 0
	for count < 8 {
		arr[count] = count + 1  // append() solo funciona con los slices.
		count += 1
	}

	refArr := otherArr[:] // Haciendo referencia a un array

	fmt.Println(arr)
	fmt.Println(otherArr)
	fmt.Println(douArr)
	fmt.Println(refArr)
	
}