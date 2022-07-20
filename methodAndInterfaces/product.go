package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}
