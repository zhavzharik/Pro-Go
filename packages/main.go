package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	//product := store.Product{
	//	Name:     "Kayak",
	//	Category: "Watersports",
	//}
	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))
	fmt.Println("Name product:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	fmt.Println("Name customer:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))

}

//go get github.com/fatih/color@v1.10.0
// для поиска пакетов
// https://pkg.go.dev
// https://github.com/golang/go/wiki/Projects
