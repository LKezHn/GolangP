package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error en returnError() function!") // errors.New() : crear un nuevo error
		return err
	} else {
		return nil
	}
}

func main(){
	err := returnError(1,2)
	if err == nil {
		fmt.Println("No hay ningun error")
	} else{
		fmt.Println(err)
	}
	
	err = returnError(5,5)
	if err == nil {
		fmt.Println("No hay ningun error")
	} else{
		fmt.Println(err)
	}

	if err.Error() == "Error en returnError() function!" { // err.Error() : convertir una variable Error a string 
		fmt.Println("!!!")
	}

}