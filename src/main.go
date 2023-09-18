package main

import (
	"fmt"

	"github.com/partherope/GoProject/td1"
)

func main() {
	A := td1.NewPoint2D(1, 2)
	B := td1.NewPoint2D(3, 4)
	RecA := td1.NewRectangle(A, B)
	fmt.Println(RecA.P1().X())
}
