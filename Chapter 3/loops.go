package main

import "fmt"

func main(){

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 9 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 10
	for { // Ciclo for emulando un ciclio while
		if i < 0 { 
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	flag := true
	for ok := true; ok; ok = flag { // for emulando un ciclo do...while
		if i > 9 {
			flag = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	anAnrray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anAnrray { // Usando range en un arreglo retorna el indice del elemento, asi como su valor
		fmt.Println("index:", i, "value:", value)
	}

}

