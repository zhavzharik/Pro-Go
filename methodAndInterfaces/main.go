package main

import "fmt"

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails()
	}
	fmt.Println()

	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}
	fmt.Println()

	products1 := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range products1.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
	fmt.Println()

	products2 := ProductList(getProducts())
	for category, total := range products2.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
	fmt.Println()

	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}

	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))
}
