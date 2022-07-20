package main

import "fmt"

type Supplier struct {
	name, city string
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}
