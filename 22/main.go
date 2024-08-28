package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y, k big.Int
	fmt.Scan(&x, &y)

	fmt.Printf("%v\n", k.Add(&x, &y))
	fmt.Printf("%v\n", k.Sub(&x, &y))
	fmt.Printf("%v\n", k.Mul(&x, &y))
	fmt.Printf("%v\n", k.Div(&x, &y))
}
