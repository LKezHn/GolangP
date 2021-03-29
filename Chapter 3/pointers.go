package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v  // Retorno su direccion
}

func main(){
	i := -5
	j := 3

	pI := &i
	pJ := &j

	fmt.Println("pI memory: ", pI )
	fmt.Println("pJ memory: ", pJ)
	fmt.Println("pI value: ", *pI)
	fmt.Println("pJ value: ", *pJ)

	*pI = 12
	fmt.Println("i:", i)
	
	*pI--
	fmt.Println("i:", i)
	
	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(6)
	fmt.Println(k)
	fmt.Println(*k)

}