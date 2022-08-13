package main

import (
	"fmt"
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	//for _, p := range Products {
	//	Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	//}

	t, err := template.ParseFiles("templates/template.html")
	if err == nil {
		t.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}
	fmt.Println()

	t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")
	if err1 == nil && err2 == nil {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t2.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error(), err2.Error())
	}
	fmt.Println()

	allTemplates, err3 := template.ParseFiles("templates/template.html", "templates/extras.html")
	if err3 == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v", err3.Error())
	}
	fmt.Println()

	allTemplates1, err4 := template.ParseGlob("templates/*.html")
	if err4 == nil {
		for _, t := range allTemplates1.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v %v", err4.Error())
	}
	fmt.Println()

	allTemplates2, err5 := template.ParseGlob("templates/*.html")
	if err5 == nil {
		selectedTemplated := allTemplates2.Lookup("template.html")
		err = Exec(selectedTemplated)
	} else {
		Printfln("Error: %v %v", err5.Error())
	}
	fmt.Println()

}
