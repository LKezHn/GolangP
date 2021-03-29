package main

import "fmt"

func main() {
	const sLiteral = "\x99\x35\x50\x44"

	fmt.Println(sLiteral)
	fmt.Printf("x : %x\n", sLiteral)
	fmt.Printf("sLiteral length : %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral)
}