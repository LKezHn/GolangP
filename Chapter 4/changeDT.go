package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"io"
)

func main(){

	arguments := os.Args
	if len(arguments) == 1{
		fmt.Println("Provide a filename to process")
		os.Exit(1)
	}

	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error to open the file %s \n", err)
		os.Exit(0)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break		
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
		}
		//* regexp.Compile() funciona igual a regexp.MustCompile() pero regexp.Compile() no lanza ningun panic.
		r1 := regexp.MustCompile(`.*\d{2}/\w+/\d{4}.*`) // Guarda una regexp compilada en una variable, lanza un panic si la regex no puede ser parseada.
		if r1.MatchString(line) {// Reporta cuando line contiene la regexp guardada en la variable r1
			match := r1.FindStringSubmatch(line) // Retorna un slice de strings que contiene el texto del match mas a la izquierda  en line 
			fmt.Println(match)
		}j
		continue
	}

}