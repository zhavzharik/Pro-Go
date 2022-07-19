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
}
