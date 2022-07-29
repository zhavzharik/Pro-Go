package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"strconv"
	"strings"
)

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"offer,string"`
}

// Creating Completely Custom JSON Encodings
func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

// Creating Completely Custom JSON Decodings
func (dp *DiscountedProduct) UnmarshalJSON(data []byte) (err error) {
	mdata := map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)
	if dp.Product == nil {
		dp.Product = &Product{}
	}
	if err == nil {
		if name, ok := mdata["Name"].(string); ok {
			dp.Name = name
		}
		if category, ok := mdata["Category"].(string); ok {
			dp.Category = category
		}
		if price, ok := mdata["Price"].(float64); ok {
			dp.Price = price
		}
		if discount, ok := mdata["Offer"].(string); ok {
			fpval, fperr := strconv.ParseFloat(discount, 64)
			if fperr == nil {
				dp.Discount = fpval
			}
		}
	}
	return
}

func main() {

	color.Cyan("Encoding JSON Data")
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival
	var writer1 strings.Builder
	encoder := json.NewEncoder(&writer1)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer1.String())
	fmt.Println()

	color.Cyan("Encoding Slices and Arrays")
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte //байтовый массив кодируется, как масив чисел в json
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0]) // срез байтов выражается в виде строки в кодировке base64
	var writer2 strings.Builder
	encoder = json.NewEncoder(&writer2)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	fmt.Print(writer2.String())
	fmt.Println()

	color.Cyan("Encoding Maps")
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	var writer3 strings.Builder
	encoder = json.NewEncoder(&writer3)
	encoder.Encode(m)
	fmt.Print(writer3.String())
	fmt.Println()

	color.Cyan("Encoding Structs")
	var writer4 strings.Builder
	encoder = json.NewEncoder(&writer4)
	encoder.Encode(Kayak)
	fmt.Print(writer4.String())
	fmt.Println()

	color.Cyan("Encoding a Struct with an Embedded Field and Struct Tags")
	var writer5 strings.Builder
	encoder = json.NewEncoder(&writer5)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	dp2 := DiscountedProduct{Discount: 10.50}
	encoder.Encode(&dp2)
	fmt.Print(writer5.String())
	fmt.Println()

	color.Cyan("Encoding Interfaces")
	var writer6 strings.Builder
	encoder = json.NewEncoder(&writer6)
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Print(writer6.String())
	fmt.Println()

	color.Cyan("Decoding Basic Data Types by providing a pointer to an empty interface")
	reader1 := strings.NewReader(`true "Hello, World!" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader1)
	decoder.UseNumber()
	for {
		var decodedVal interface{}         // the Decoder is able to select the appropriate Go data type for JSON values,
		err := decoder.Decode(&decodedVal) // and this is achieved by providing a pointer to an empty interface as the argument to the Decode method
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
	fmt.Println()

	color.Cyan("Decoding Basic Data Types specifying types for decoding")
	reader2 := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var intval int
	vals2 := []interface{}{&bval, &sval, &fpval, &intval}
	decoder = json.NewDecoder(reader2)
	for i := 0; i < len(vals2); i++ {
		err := decoder.Decode(vals2[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", intval, intval)
	fmt.Println()

	color.Cyan("Decoding Arrays by providing a pointer to an empty interface")
	reader3 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	vals3 := []interface{}{}
	decoder = json.NewDecoder(reader3)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals3 = append(vals3, decodedVal)
	}
	for _, val := range vals3 {
		Printfln("Decoded (%T): %v", val, val)
	}
	fmt.Println()

	color.Cyan("Decoding Arrays specifying types for decoding")
	reader4 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	ints := []int{}
	mixed := []interface{}{}
	vals4 := []interface{}{&ints, &mixed}
	decoder = json.NewDecoder(reader4)
	for i := 0; i < len(vals4); i++ {
		err := decoder.Decode(vals4[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)
	fmt.Println()

	color.Cyan("Decoding Maps")
	reader5 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m5 := map[string]interface{}{}
	decoder = json.NewDecoder(reader5)
	err := decoder.Decode(&m5)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Decoded Map: (%T), %v", m5, m5)
		for k, v := range m5 {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
	fmt.Println()

	color.Cyan("Decoding Maps specifying type for value")
	reader6 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m6 := map[string]float64{}
	decoder = json.NewDecoder(reader6)
	err = decoder.Decode(&m6)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Decoded Map: (%T), %v", m6, m6)
		for k, v := range m6 {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
	fmt.Println()

	color.Cyan("Decoding Structs")
	reader7 := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279}
		{"Name":"Lifejacket","Category":"Watersports"}
		{"name":"Canoe","category":"Watersports", "price": 100, "inStock": true}`)
	decoder = json.NewDecoder(reader7)
	decoder.DisallowUnknownFields()
	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		}
	}
	fmt.Println()

	reader8 := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer":"10"}`)
	decoder = json.NewDecoder(reader8)
	for {
		var val DiscountedProduct
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
				val.Name, val.Category, val.Price, val.Discount)
		}
	}

}
