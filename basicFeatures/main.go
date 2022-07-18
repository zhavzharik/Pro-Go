package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	const price, tax float32 = 275.00, 27.5
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity * (price + tax))
	fmt.Println("In stock:", inStock)
	//fmt.Println("Value:", rand.Int())
}