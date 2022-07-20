package main

import "fmt"

func printPrice1(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

func printPrice2(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

func printPrice3(string, float64, float64) {
	fmt.Println("No parameters")
}

func printSuppliers1(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func printSuppliers2(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func swapValues1(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

func swapValues2(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

func swapValues3(first, second int) (int, int) {
	return second, first
}

func calcTax1(price float64) float64 {
	if price > 100 {
		return price * 0.2
	}
	return -1
}

func calcTax2(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice1(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax2(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func calcTotalPrice2(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}
	return
}

func calcTotalPrice3(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	printPrice1("Kayak", 275, 0.2)
	printPrice1("KifeJacket", 48.95, 0.2)
	printPrice1("Soccer Ball", 19.50, 0.15)
	fmt.Println("")

	printPrice2("Kayak", 275, 0.2)
	printPrice2("KifeJacket", 48.95, 0.2)
	printPrice2("Soccer Ball", 19.50, 0.15)
	fmt.Println("")

	printPrice3("Kayak", 275, 0.2)
	fmt.Println("")

	printSuppliers1("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers1("Lifejacket", []string{"Sail Safe Co"})
	fmt.Println("")

	printSuppliers2("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers2("Lifejacket", "Sail Safe Co")
	printSuppliers2("Soccer Ball")
	fmt.Println("")

	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers2("Kayak", names...)
	printSuppliers2("Lifejacket", "Sail Safe Co")
	printSuppliers2("Soccer Ball")
	fmt.Println("")

	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues1(val1, val2)
	fmt.Println("After calling function", val1, val2)
	fmt.Println("")

	fmt.Println("Before calling function", val1, val2)
	swapValues2(&val1, &val2)
	fmt.Println("After calling function", val1, val2)
	fmt.Println("")

	products := map[string]float64{
		"Kayak ":     275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Println("Product:", product, "Price:", priceWithTax)
	}

	val1, val2 = 10, 20
	fmt.Println("Before calling function", val1, val2)
	val1, val2 = swapValues3(val1, val2)
	fmt.Println("After calling function", val1, val2)
	fmt.Println("")

	for product, price := range products {
		tax := calcTax1(price)
		if tax != -1 {
			fmt.Println("Product:", product, "Tax:", tax)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}
	fmt.Println("")

	for product, price := range products {
		if taxAmount, taxDue := calcTax2(price); taxDue {
			fmt.Println("Product:", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}
	fmt.Println("")

	total1, tax1 := calcTotalPrice1(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice1(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
	fmt.Println("")

	_, total := calcTotalPrice2(products)
	fmt.Println("Total:", total)
	fmt.Println("")

	_, total3 := calcTotalPrice3(products)
	fmt.Println("Total:", total3)
}
