package main

import (
	"fmt"
)

// Map en Go -> Diccionario en Python

func main(){
	iMap := make(map[string]int) // Declarando un map con la función make().
	
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Printf("iMap: %v \n", iMap)

	otherMap := map[string] int {
		"k1": 12,
		"k2": 13,
	} 	
	fmt.Println(otherMap["k1"])
	fmt.Println(otherMap["k1"])
	fmt.Println(otherMap["k1"])
	delete(otherMap, "a") //Borrando el elemnto con key "a" en el map "OhterMap"
	delete(otherMap, "a") // Llamar multiples veces la funcion delete() no genera ninguna advertencia.
	fmt.Println(otherMap)

	_, ok := iMap["keyExist"] //  Obteniendo el key y el value del elemento con key "keyExist".

	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Doesn't exist")
	}

	for key, value := range iMap {
		fmt.Printf("key: %v, value: %v \n", key, value)
	} 

	aMap := map[string]int{}
	aMap["one"] = 1

	fmt.Println(aMap)

	/*
		El siguiente codigo da error ya que asignamos el valor de nil a la variable aMap por lo que ya no es de tipo map.
	aMap = nil
	fmt.Println(aMap)
	aMap["one"] = 1
	*/
}