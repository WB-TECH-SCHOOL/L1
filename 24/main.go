package main

import (
	"24/point"
	"fmt"
)

func main() {
	p1 := point.NewPoint(9, 6)
	p2 := point.NewPoint(5, 2)
	fmt.Println(p1.GetDistant(p2))
}
