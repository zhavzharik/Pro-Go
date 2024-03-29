package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {

	// Working with arrays and slices
	var names [3]string

	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	names1 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names1)
	fmt.Println("")

	otherArray := names1
	names1[0] = "Canoe"
	fmt.Println("names:", names1)
	fmt.Println("otherArray:", otherArray)
	fmt.Println("")

	otherArray1 := &names1
	names1[0] = "Canoe"
	fmt.Println("names:", names1)
	fmt.Println("otherArray:", *otherArray1)

	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := names == moreNames
	fmt.Println("Comparison:", same)
	fmt.Println("")

	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("")

	names2 := make([]string, 3)
	names2[0] = "Kayak"
	names2[1] = "Lifejacket"
	names2[2] = "Paddle"
	fmt.Println(names2)

	names3 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names3)
	fmt.Println("")

	appendedNames := append(names3, "Hat", "Gloves")
	fmt.Println("names:", names3)
	fmt.Println("appendedNames:", appendedNames)
	fmt.Println("")

	names4 := make([]string, 3, 6)
	names4[0] = "Kayak"
	names4[1] = "Lifejacket"
	names4[2] = "Paddle"
	fmt.Println("len:", len(names4))
	fmt.Println("cap:", cap(names4))
	fmt.Println("")

	appendedNames1 := append(names4, "Hat", "Gloves")
	names4[0] = "Canoe"
	fmt.Println("names:", names4)
	fmt.Println("appendedNames:", appendedNames1)
	fmt.Println("len appendedNames:", len(appendedNames1))
	fmt.Println("cap appendedNames:", cap(appendedNames1))
	fmt.Println("")

	moreNames1 := []string{"Hat Gloves"}
	appendedNames2 := append(names4, moreNames1...)
	fmt.Println("appendedNames:", appendedNames2)
	fmt.Println("len appendedNames:", len(appendedNames2))
	fmt.Println("cap appendedNames:", cap(appendedNames2))
	fmt.Println("")

	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println("")

	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println("")

	someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println("")

	someNames[0] = "Jacket"
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println("")

	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Println("products:", products)
	fmt.Println("products len:", len(products), "cap:", cap(products))
	someNames = products[1:3:3]
	allNames = products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	//products[1] = "Canoe"
	//fmt.Println("someNames:", someNames)
	//fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	allNames = append(allNames, "Glasses")
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	products[1] = "Jacket"
	fmt.Println("products:", products)
	fmt.Println("products len:", len(products), "cap:", cap(products))
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
	fmt.Println("")

	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames = products[1:]
	someNames = allNames[1:3]
	fmt.Println("allNames", allNames)
	fmt.Println("someNames:", someNames)
	//allNames = append(allNames, "Gloves")
	//fmt.Println("allNames", allNames)
	allNames[1] = "Canoe"
	fmt.Println("allNames", allNames)
	fmt.Println("someNames:", someNames)
	fmt.Println("")

	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames = products[1:]
	someNames = make([]string, 2)

	copy(someNames, allNames)
	fmt.Println("allNames", allNames)
	fmt.Println("someNames:", someNames)

	var someNames1 []string
	copy(someNames1, allNames)
	fmt.Println("allNames", allNames)
	fmt.Println("someNames:", someNames1)

	someNames = []string{"Boots", "Canoe"}
	copy(someNames[1:], allNames[2:3])
	fmt.Println("allNames", allNames)
	fmt.Println("someNames:", someNames)
	fmt.Println("")

	products1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products1, replacementProducts)
	fmt.Println("products:", products1)

	products1 = []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	copy(products1[0:1], replacementProducts)
	fmt.Println("products:", products1)
	fmt.Println("")

	products = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Println("products:", products)
	deleted := append(products[:2], products[3:]...)
	fmt.Println("deleted:", deleted)
	fmt.Println("")

	products1 = []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products1[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("")

	sort.Strings(products1)
	for index, value := range products1 {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("products:", products1)
	fmt.Println("")

	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	//fmt.Println("Equal:", p1 == p2) // так нельзя сравнить срезы
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))

	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println(array)
	fmt.Println("")

	//Working with maps

	products2 := make(map[string]float64, 10)
	products2["Kayak"] = 279
	products2["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(products2))
	fmt.Println("Price:", products2["Kayak"])
	fmt.Println("Price:", products2["Hat"])
	fmt.Println("")

	products3 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	fmt.Println("Map size:", len(products3))
	fmt.Println("Price:", products3["Kayak"])
	value, ok := products3["Hat"] // ok = true, если карта содержит указанный ключ
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
	delete(products3, "Hat")
	if value, ok := products3["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
	fmt.Println("")

	products3 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	for key, value := range products3 {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("")

	products3 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	keys := make([]string, 0, len(products3))
	for key, _ := range products3 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products3[key])
	}
	fmt.Println("")

	// Understanding the dual nature of strings
	var price []rune = []rune("€48.95")

	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
	fmt.Println("")

	var price1 = "€48.95"
	for index, char := range price1 {
		fmt.Println(index, char, string(char))
	}
	fmt.Println("")
	for index, char := range []byte(price1) {
		fmt.Println(index, char)
	}
}
