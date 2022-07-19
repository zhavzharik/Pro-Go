package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
	fmt.Println("")

	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))

	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)
	fmt.Println("")

	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)
	fmt.Println("")

	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)
	fmt.Println("")

	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
	fmt.Println("")

	newSecond := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(newSecond == third)
	fmt.Println(newSecond == beta)

	fmt.Println(*newSecond == *third)
	fmt.Println(*newSecond == *beta)
	fmt.Println("")

	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
	fmt.Println("")

	kayak := 275
	soccerBall := 19.50
	total1 := float64(kayak) + soccerBall
	total2 := kayak + int(math.Round(soccerBall))
	fmt.Println(total1)
	fmt.Println(total2)
	fmt.Println("")

	value27 := 27.1
	fmt.Println(math.Ceil(value27))
	fmt.Println(math.Floor(value27))
	fmt.Println(math.RoundToEven(value27))
	fmt.Println("")

	val1 := "0"
	val2 := "false"
	val3 := "not true"
	if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
		fmt.Println("Parsed Bool 1 value:", bool1)
	} else {
		fmt.Println("Cannot parse Bool 1", val1)
	}
	//fmt.Println("Bool 1", bool1, b1err)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)
	fmt.Println("")

	val100 := "100"
	int1, int1err := strconv.ParseInt(val100, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}

	val500 := "500"
	int5, int5err := strconv.ParseInt(val500, 0, 8)
	if int5err == nil {
		fmt.Println("Parsed value:", int5)
	} else {
		fmt.Println("Cannot parse", val500, int5err)
	}
	fmt.Println("")

	value100 := "100"
	int100, int100err := strconv.ParseInt(value100, 2, 8)
	if int100err == nil {
		smallInt100 := int8(int100)
		fmt.Println("Parsed value:", smallInt100)
	} else {
		fmt.Println("Cannot parse", value100, int100err)
	}

	val0 := "0b1100100"
	int0, int0err := strconv.ParseInt(val0, 0, 8)
	if int0err == nil {
		smallInt := int8(int0)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}

	int10, int10err := strconv.ParseInt(value100, 10, 0)
	if int10err == nil {
		smallInt10 := int8(int10)
		fmt.Println("Parsed value:", smallInt10)
	} else {
		fmt.Println("Cannot parse", value100, int10err)
	}

	intA, intAerr := strconv.Atoi(value100)
	if intAerr == nil {
		var intResult int = intA
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", value100, intAerr)
	}
	fmt.Println("")

	fl1 := "48.95"
	float1, float1err := strconv.ParseFloat(fl1, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", fl1, float1err)
	}

	fl2 := "4.895e+01"
	float2, float2err := strconv.ParseFloat(fl2, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float2)
	} else {
		fmt.Println("Cannot parse", fl2, float2err)
	}
	fmt.Println("")

	v1 := true
	v2 := false
	str1 := strconv.FormatBool(v1)
	str2 := strconv.FormatBool(v2)
	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)

	v := 275
	base10String := strconv.FormatInt(int64(v), 10)
	b10String := strconv.Itoa(v)
	base2String := strconv.FormatInt(int64(v), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 10: " + b10String)
	fmt.Println("Base 2: " + base2String)

	x1 := 49.95
	Fstring := strconv.FormatFloat(x1, 'f', 2, 64)
	Estring := strconv.FormatFloat(x1, 'e', -1, 64)
	Gstring := strconv.FormatFloat(x1, 'g', 1, 64)
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
	fmt.Println("Format G: " + Gstring)
}
