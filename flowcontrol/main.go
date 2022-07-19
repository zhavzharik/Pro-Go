package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 375.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else if kayakPrice > 200 && kayakPrice < 300 {
		scopedVar := "Price is between 200 and 300"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}
	fmt.Println("")

	priceString := "275"

	if kPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price: ", kPrice)
	} else {
		fmt.Println("Error:", err)
	}
	fmt.Println("")

	//counter := 0
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
		//counter++
		//if counter > 3 {
		//	break
		//}
	}
	fmt.Println("")
	for counter := 0; true; counter++ {
		fmt.Println("Counter:", counter)
		if counter > 3 {
			break
		}
	}
	fmt.Println("")
	for counter := 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
	fmt.Println("")

	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}
	for index := range product {
		fmt.Println("Index:", index)
	}
	for _, character := range product {
		fmt.Println("Character:", string(character))
	}

	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}

	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
	fmt.Println("")

	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
	fmt.Println("")

	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}
	fmt.Println("")

	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}
	fmt.Println("")

	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
	fmt.Println("")

	counter := 0
	// метки определяются именем, за которым следует двоеточие
	// ключевое слово goto используется для перехода к метке
target:
	fmt.Println("Counter", counter) // кодовый оператор
	counter++
	if counter < 5 {
		goto target
	}

}
