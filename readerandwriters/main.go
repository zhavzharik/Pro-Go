package main

import (
	"bufio"
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

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
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
	fmt.Println()

	color.Cyan("Using TeeReader(reader, writer)")
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())
	fmt.Println()

	color.Cyan("Using LimitReader(reader, n)")
	r1 = strings.NewReader("Kayak")
	r2 = strings.NewReader("Lifejacket")
	r3 = strings.NewReader("Canoe")
	concatReader = io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
	fmt.Println()

	color.Cyan("CustomReader")
	text := "It was a boat. A small boat."
	var readerX io.Reader = NewCustomReader(strings.NewReader(text))
	var writerX strings.Builder
	slice := make([]byte, 5)
	for {
		count, err := readerX.Read(slice)
		if count > 0 {
			writerX.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writerX.String())
	fmt.Println()

	color.Cyan("Using Buffered Reader")
	var readerB io.Reader = NewCustomReader(strings.NewReader(text))
	var writerB strings.Builder
	buffered := bufio.NewReader(readerB)
	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writerB.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writerB.String())
	fmt.Println()

	color.Cyan("CustomWriter")
	var builderY strings.Builder
	var writerY = NewCustomWriter(&builderY)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writerY.Write([]byte(text[i:]))
			break
		}
		writerY.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builderY.String())
	fmt.Println()

	color.Cyan("Using Buffered Writer")
	var builderBW strings.Builder
	var writerBW = bufio.NewWriterSize(NewCustomWriter(&builderBW), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writerBW.Write([]byte(text[i:]))
			writerBW.Flush()
			break
		}
		writerBW.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builderBW.String())
	fmt.Println()

	color.Cyan("Scanning Values from a Reader")
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}

	color.Cyan("Scanning Single Value from a Reader")
	reader = strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
	fmt.Println()

	color.Cyan("Writing Formatted Strings to a Writer")
	var writerF strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writerF, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writerF.String())
	fmt.Println()

	color.Cyan("Using a Replacer with a Writer")
	subs := []string{"boat", "kayak", "small", "huge"}
	fmt.Println(text)
	fmt.Println("List of old, new string pairs:", subs)
	var writerR strings.Builder
	writeReplaced(&writerR, text, subs...)
	fmt.Println(writerR.String())
}
