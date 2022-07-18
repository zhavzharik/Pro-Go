package main

import (
	"fmt"
	"sort"
	// "math/rand"
)

func main() {
	// var price, tax, price3 float32 = 275.0, 27.5, 300.5
	// var price1 = 377.0
	// var price2 = price1
	// price4, price3, tax := 1000, 304.00, 30.4
	// const quantity, inStock = 2, true
	// x, y, z := 100, 0.5, true
	// var _ = "Alice"
	// fmt.Println("Total1:", 2 * quantity * (price + tax))
	// fmt.Println("In stock:", inStock)
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)
	// fmt.Println(price1)
	// fmt.Println(price2)
	// fmt.Println(price4)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println("Total2:", 2 * quantity * (price3 + tax))
	// fmt.Println("Value:", rand.Int())

	first := 100
	a := first
	second := &first
	third := &second
	first++
	*second++
	var myNewPointer *int
	myNewPointer = second
	*myNewPointer++
	fmt.Println("First:", first)
	fmt.Println("A:", a)
	fmt.Println("Second:", *second)
	fmt.Println("My new pointer:", *myNewPointer)
	fmt.Println("Third:", **third)
	var ptr *int
	fmt.Println("Not defined:", ptr)
	ptr = &first
	fmt.Println("ptr = &first:", ptr)

	names := [3]string {"Alice", "Charlie", "Bob"}
	secondName := names[1]
	secondPosition := &names[1]
	fmt.Println(secondName)
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println(*secondPosition)	
}