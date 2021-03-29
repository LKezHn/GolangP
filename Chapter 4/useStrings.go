package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	phrase := "Hello World"
	upper := s.ToUpper(phrase)
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower(phrase))
	f("%s\n", s.Title("a title here "))
	//Verificar si dos strings son iguales sin importar la diferencia entre letras mayusculas y minusculas
	f("EqualFold %v\n", s.EqualFold("Mihalis", "MIHAlis")) 
	f("EqualFold %v\n", s.EqualFold("Mihalis", "MIHAli"))

	// Compare Prefix
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi")) 
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	
	// Compare Suffix
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is")) 
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index: %v\n", s.Index("Mihalis", "ha")) // Retorna el indice donde donde se encuentra la letra o silaba dada 
	f("Index: %v\n", s.Index("Mihalis", "Ha")) // Si no la encunetra retorna -1
	f("Count: %v\n", s.Count("Mihalis","i"))
	f("Count: %v\n", s.Count("Mihalis","I"))
	f("Repeat: %v\n", s.Repeat("ab", 6))

	f("Trim: %s", s.TrimSpace("\tHello World.\n")) 
	/*
		Tambien existen las funciones TrimLeft() y TrimRight()
	*/

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS")) // returns 1
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis")) // returns 0
	f("Compare: %v\n", s.Compare("MIHALIS", "Mihalis")) // returns -1

	f("%s\n", s.Split("hello,world,here",","))

	f("%s\n", s.Replace("hola mundo","","_",-1))
	f("%s\n", s.Replace("hola mundo","","_",1))
	f("%s\n", s.Replace("hola mundo","","_",5))

	test := []string{"Hola","mundo","xdxd"}
	myString := s.Join(test," ")
	fmt.Println(myString)

	f("SplitAfter %s\n", s.SplitAfter("Hola--Mundo--aiuda","--")) // return a slice
}