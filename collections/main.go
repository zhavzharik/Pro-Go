package main

import (
	"fmt"
	"reflect"
	"sort"
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
}
