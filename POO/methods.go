package main

import ( 
	"math" // To import PI
	"fmt"
) 

type Circle struct{
	radius float64
}

type Rectangle struct{
	width, height float64
}

// ! Method Prototype
// func (r Receiver) functionName(params) (results)

// Area de un circulo
func (c Circle) Area() float64 {
	return c.radius * c.radius + math.Pi
}

// Area de un rectnagulo
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {

	c1 := Circle{10}
	c2 := Circle{25}
	// No da error porque las funciones aunque se llaman igual
	// sus Receivers son distintos
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	fmt.Printf("Area of c1: %v\n", c1.Area())
	fmt.Printf("Area of c2: %v\n", c2.Area())
	fmt.Printf("Area of r1: %v\n", r1.Area())
	fmt.Printf("Area of r2: %v\n", r2.Area())

}