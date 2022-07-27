package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {

	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Println()
	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
	fmt.Println()

	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())

	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	Printfln("Value with fields: %+v", Kayak)

	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)

	nb := 279.00
	Printfln("Decimalless with exponent: %b", nb)
	Printfln("Decimalless without exponent: >>%8.2f<<", nb)
	Printfln("Decimal with exponent: %e", nb)
	Printfln("Decimal without exponent: %f", nb)
	Printfln("Hexadecimal: %x, %X", nb, nb)
	Printfln("Sign: >>%+.2f<<", nb)
	Printfln("Zeros for Padding: >>%010.2f<<", nb)
	Printfln("Right Padding: >>%-8.2f<<", nb)

	name = "Kayak"
	Printfln("String: %s", name)
	Printfln("Character: %c", []rune(name)[0])
	Printfln("Unicode: %U", []rune(name)[0])

	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)
	Printfln("Pointer: %p", &name)

	var nameX string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scanln(&nameX, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", nameX, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}

	source := "Lifejacket Watersports 48.95"
	n, err = fmt.Sscan(source, &nameX, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", nameX, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
	fmt.Println()

	source = "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err = fmt.Sscanf(source, template, &nameX, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", nameX, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}

}
