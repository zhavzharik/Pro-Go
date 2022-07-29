package main

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"strings"
)

func processDataOnlyRead(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func processDataCopy(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	Printfln("Product: %v, Price : %v", Kayak.Name, Kayak.Price)
	fmt.Println()

	color.Cyan("Using Read(byteSlice)")
	r := strings.NewReader("Kayak")
	processDataOnlyRead(r)
	fmt.Println()

	color.Cyan("Using Write(byteSlice)")
	r = strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
	fmt.Println()

	color.Cyan("Using Copy(w, r)")
	r = strings.NewReader("Kayak")
	var builderForCopy strings.Builder
	processDataCopy(r, &builderForCopy)
	Printfln("String builder contents: %s", builderForCopy.String())
	fmt.Println()

	color.Cyan("Using Pipe()")
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
	fmt.Println()

	color.Cyan("Using MultiReader(...readers)")
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
	fmt.Println()

	color.Cyan("Using MultiWriter(...writers)")
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

	color.Cyan("Using TeeReader(reader, writer)")
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())

	color.Cyan("Using LimitReader(reader, n)")
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)

}
