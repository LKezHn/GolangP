package main

import (
	"fmt"
)

const (
	WHITE = iota // Generando numeros automaticamente
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct{
	width, height, depth float64
	color Color
}

type Color byte // Customized
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// Se usa puntero porque se requiere cambiar el valor de la caja original.
// Si no se usa el puntero se cambiará el valor de una copia de la caja.
func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}

	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) ToString() string {
	colors := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return colors[c]
}

func main() {
	boxes := BoxList{
		Box{4,4,4, RED},
		Box{10,10,1, YELLOW},
		Box{1,1,20, BLACK},
		Box{10,10,1, BLUE},
		Box{10,30,1, WHITE},
		Box{20,20,20, YELLOW},
	}

	fmt.Printf("Boxes in BoxList. %v\n", len(boxes))
	fmt.Printf("Volume of the first: %v cm³\n", boxes[0].Volume())
	fmt.Printf("Color of the last: %v\n", boxes[len(boxes) - 1].color.ToString())
	fmt.Printf("The biggest is: %v\n", boxes.BiggestColor().ToString())
	
	boxes.PaintItBlack()
	
	fmt.Printf("Color of the second: %v\n", boxes[1].color.ToString())
	fmt.Printf("The biggest is: %v\n", boxes.BiggestColor().ToString())
	
}