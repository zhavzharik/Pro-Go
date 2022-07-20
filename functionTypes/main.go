package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice1(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator1(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

type calcFunc func(float64) float64

func printPrice2(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator2(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func selectCalculator3(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(float64) float64 {
		return price
	}
	return withoutTax
}

func selectCalculator4(price float64) calcFunc {
	if price > 100 {
		return func(price float64) float64 {
			return price + (price * 0.2)
		}
	}
	return func(float64) float64 {
		return price
	}
}

func priceCalcFactory1(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

var prizeGiveaway = false

func priceCalcFactory2(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func priceCalcFactory3(threshold, rate float64) calcFunc {
	fixedPrizeGiveaway := prizeGiveaway
	return func(price float64) float64 {
		if fixedPrizeGiveaway {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func priceCalcFactory4(threshold, rate float64, zeroPrices bool) calcFunc {
	return func(price float64) float64 {
		if zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func priceCalcFactory5(threshold, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
	fmt.Println("")

	for product, price := range products {
		if price > 100 {
			printPrice1(product, price, calcWithTax)
		} else {
			printPrice1(product, price, calcWithoutTax)
		}
	}
	fmt.Println("")

	for product, price := range products {
		printPrice1(product, price, selectCalculator1(price))
	}
	fmt.Println("")

	for product, price := range products {
		if price > 100 {
			printPrice2(product, price, calcWithTax)
		} else {
			printPrice2(product, price, calcWithoutTax)
		}
	}
	fmt.Println("")

	for product, price := range products {
		printPrice2(product, price, selectCalculator2(price))
	}
	fmt.Println("")

	for product, price := range products {
		printPrice2(product, price, selectCalculator3(price))
	}
	fmt.Println("")

	for product, price := range products {
		printPrice2(product, price, selectCalculator4(price))
	}
	fmt.Println("")

	for product, price := range products {
		printPrice2(product, price, func(price float64) float64 {
			return price + (price * 0.2)
		})
	}
	fmt.Println("")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	calc := func(price float64) float64 {
		if price > 100 {
			return price + (price * 0.2)
		}
		return price
	}
	for product, price := range watersportsProducts {
		printPrice2(product, price, calc)
	}

	calc = func(price float64) float64 {
		if price > 50 {
			return price + (price * 0.1)
		}
		return price
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, calc)
	}
	fmt.Println("")

	waterCalc := priceCalcFactory1(100, 0.2)
	soccerCalc := priceCalcFactory1(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice2(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, soccerCalc)
	}
	fmt.Println("")

	prizeGiveaway = false
	waterCalc = priceCalcFactory2(100, 0.2)
	prizeGiveaway = true
	soccerCalc = priceCalcFactory2(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice2(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, soccerCalc)
	}
	fmt.Println("")

	prizeGiveaway = false
	waterCalc = priceCalcFactory3(100, 0.2)
	prizeGiveaway = true
	soccerCalc = priceCalcFactory3(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice2(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, soccerCalc)
	}
	fmt.Println("")

	prizeGiveaway = false
	waterCalc = priceCalcFactory4(100, 0.2, prizeGiveaway)
	prizeGiveaway = true
	soccerCalc = priceCalcFactory4(50, 0.1, prizeGiveaway)

	for product, price := range watersportsProducts {
		printPrice2(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, soccerCalc)
	}
	fmt.Println("")

	prizeGiveaway = false
	waterCalc = priceCalcFactory5(100, 0.2, &prizeGiveaway)
	prizeGiveaway = true
	soccerCalc = priceCalcFactory5(50, 0.1, &prizeGiveaway)

	for product, price := range watersportsProducts {
		printPrice2(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice2(product, price, soccerCalc)
	}
}
