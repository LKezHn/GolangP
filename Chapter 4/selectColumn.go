package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"io"
)

func main(){
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("usage: selectColumn column <file1> [<file2>... [fileN]]")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(arguments[1]) // Convirtiendo el argumento "column" en int.
	if err != nil {
		fmt.Println("Column argument value is not an int: ", temp)
		return 
	}

	column := temp
	if column < 0 {
		fmt.Println("Column number is invalid: ", column)
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Eror opening the file %v. \n", err)
			continue // Continua ya que puede tener mas de un archivo para leer.
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n') // Lee cada linea hasta encontrar el delimitador enviado en la funcion.

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading the file %s", err)
			}

			data := strings.Fields(line) // strings.Fields(line) in Go -> line.split(" ") in Python.
			if len(data) >= column {
				fmt.Println((data[column - 1]))
			}
		}
	}
}