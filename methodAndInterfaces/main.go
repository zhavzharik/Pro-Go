package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}

//func (p *Product) getName() string {
//	return p.name
//}
//
//func (p *Product) getCost(_ bool) float64 {
//	return p.price
//}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

func (s ServiceWithFeatures) getName() string {
	return s.description
}

func (s ServiceWithFeatures) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

type Person struct {
	name, city string
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case ServiceWithFeatures:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}

func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case ServiceWithFeatures:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails()
	}
	fmt.Println()

	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}
	fmt.Println()

	products1 := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range products1.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
	fmt.Println()

	products2 := ProductList(getProducts())
	for category, total := range products2.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
	fmt.Println()

	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}

	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))
	fmt.Println()

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(expenses))
	fmt.Println()

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))
	fmt.Println()

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
	fmt.Println()

	product = Product{"Kayak", "Watersports", 275}
	var expense1 Expense = &product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense1.getCost(false))
	fmt.Println()

	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)
	fmt.Println()

	// var e5 Expense = ServiceWithFeatures { description: "Boat Cover" }
	// var e6 Expense = ServiceWithFeatures { description: "Boat Cover" }
	// fmt.Println("e5 == e6", e5 == e6) // срезы несопоставимы

	expenses1 := []Expense{
		ServiceWithFeatures{"Boat Cover", 12, 89.50, []string{}},
		ServiceWithFeatures{"Paddle Protect", 12, 8, []string{}},
	}
	for _, expense := range expenses1 {
		// утверждение типа, можно применять только к интерфейсам,
		// используется для того, чтобы сообщить компилятору, что значение интерфейса имеет определенный динамический тип
		s := expense.(ServiceWithFeatures)
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
	fmt.Println()

	expenses2 := []Expense{
		ServiceWithFeatures{"Boat Cover", 12, 89.50, []string{}},
		ServiceWithFeatures{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses2 {
		if s, ok := expense.(ServiceWithFeatures); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}
	fmt.Println()

	for _, expense := range expenses2 {
		switch value := expense.(type) {
		case ServiceWithFeatures:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}
	fmt.Println()

	// использование пустого интерфейса
	var expense2 Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense2,
		Product{"Lifejacket", "Watersports", 48.95},
		ServiceWithFeatures{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case ServiceWithFeatures:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
	fmt.Println()

	for _, item := range data {
		processItem(item)
	}
	fmt.Println()

	processItems(data...)
}
