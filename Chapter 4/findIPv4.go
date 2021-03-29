package main

import (
	"regexp"
	"fmt"
	"os"
	"io"
	"net"
	"bufio"
	"path/filepath"
)

func findIP(in string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])" 
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(in) //Retorna el fragmento o palabra donde matchea la expresion regular.
}

func main() {
	arguments := os.Args
	if len(arguments) < 2{
		fmt.Printf("Usage %s logFile \n", filepath.Base(arguments[0])) // filepath.Base(path) Retorna el último elemento del path dado 
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)

		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error to read the file %s", err)
			}

			ip := findIP(line)
			trial := net.ParseIP(ip) // Comprueba si la varaible ip contiene la estructura correcta de una direccion IPv4.
			if trial.To4() == nil { // Si trial.TO4() es nil significa que la IP no es de tipo IPv4.
				continue
			}
			fmt.Println(ip)
		}
	}
}