package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	name, category string
	price          float64
}

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func calcTax1(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func calcTax2(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

type Product1 struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct1(name, category string, price float64, supplier *Supplier) *Product1 {
	return &Product1{name, category, price - 10, supplier}
}

func copyProduct1(product *Product1) Product1 {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func main() {

	//type Product struct {
	//	name, category string
	//	price          float64
	//}
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		//price:    275,
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
	fmt.Println("")

	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)
	fmt.Println("")

	//взаимозаменяемые подходы
	//var lifejacket = new(Product)
	//var lifejacket = &Product{}

	var kayak1 = Product{"Kayak", "Watersports", 275.00}
	fmt.Println("Name:", kayak1.name)
	fmt.Println("Category:", kayak1.category)
	fmt.Println("Price:", kayak1.price)
	fmt.Println("")

	type StockLevel1 struct {
		Product
		count int
	}

	stockItem1 := StockLevel1{
		Product: Product{"Kayak", "Watersports", 275.00},
		count:   100,
	}
	fmt.Println("Name:", stockItem1.Product.name)
	fmt.Println("Count:", stockItem1.count)

	fmt.Println(fmt.Sprint("Name: ", stockItem1.Product.name))
	fmt.Println("")

	type StockLevel2 struct {
		Product
		Alternate Product
		count     int
	}

	stockItem2 := StockLevel2{
		Product:   Product{"Kayak", "Watersports", 275.00},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("Name:", stockItem2.Product.name)
	fmt.Println("Alt Name:", stockItem2.Alternate.name)
	fmt.Println("")

	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
	fmt.Println("")

	//type Product1 struct {
	//	name, category string
	//	price          float64
	//	otherNames     []string
	//}
	//
	//pr1 := Product1{name: "Kayak", category: "Watersports", price: 275.00}
	//pr2 := Product1{name: "Kayak", category: "Watersports", price: 275.00}
	//pr3 := Product1{name: "Kayak", category: "Boats", price: 275.00}
	//fmt.Println("p1 == p2:", pr1 == pr2)
	//fmt.Println("p1 == p3:", pr1 == pr3)

	type Item struct {
		name     string
		category string
		price    float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item:", prod == Product(item))
	fmt.Println("")

	item = Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod)
	writeName(item)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())

	array := [1]StockLevel2{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)
	slice := []StockLevel2{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel2{
		"kayak": {Product: Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)
	fmt.Println("")

	p1 = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275}
	p2 = p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)

	p3 = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275}
	p4 := &p3
	p3.name = "Original Kayak"
	fmt.Println("P1:", p3.name)
	fmt.Println("P2:", (*p4).name)
	fmt.Println("")

	kayak3 := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax1(kayak3)
	fmt.Println("Name:", kayak3.name, "Category:", kayak3.category, "Price:", kayak3.price)

	kayak4 := calcTax2(&Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})
	fmt.Println("Name:", kayak4.name, "Category:", kayak4.category, "Price:", kayak4.price)
	fmt.Println("")

	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
	}

	acme := &Supplier{"Acme Co", "New York"}
	products1 := [2]*Product1{
		newProduct1("Kayak", "Watersports", 275, acme),
		newProduct1("Hat", "Skiing", 42.50, acme),
	}
	fmt.Println("")

	for _, p1 := range products1 {
		fmt.Println("Name:", p1.name, "Supplier:", p1.Supplier.name, p1.Supplier.city)
	}
	fmt.Println("")

	prod1 := newProduct1("Kayak", "Watersports", 275, acme)
	prod2 := *prod1

	prod1.name = "Original Kayak"
	prod1.Supplier.name = "BoatCo"

	for _, prod := range []Product1{*prod1, prod2} {
		fmt.Println("Name:", prod.name, "Supplier:", prod.Supplier.name, prod.Supplier.city)
	}
	fmt.Println("")

	acme = &Supplier{"Acme Co", "New York"}
	prod3 := newProduct1("Kayak", "Watersports", 275, acme)
	prod4 := copyProduct1(prod3)

	prod3.name = "Original Kayak"
	prod3.Supplier.name = "BoatCo"

	for _, prod := range []Product1{*prod3, prod4} {
		fmt.Println("Name:", prod.name, "Supplier:", prod.Supplier.name, prod.Supplier.city)
	}
	fmt.Println("")

	var prodZero Product
	var prodZeroPtr *Product

	fmt.Println("Value:", prodZero.name, prodZero.category, prodZero.price)
	fmt.Println("Pointer:", prodZeroPtr)

	var prodZero1 Product1 = Product1{Supplier: &Supplier{}}
	var prodZeroPtr1 *Product1

	fmt.Println("Value:", prodZero1.name, prodZero1.category, prodZero1.price, prodZero1.Supplier.name)
	fmt.Println("Pointer:", prodZeroPtr1)
}
