package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	total := 0.0
	for _, p := range Products {
		//Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
		total += p.Price
	}

	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)
	//err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		//fmt.Println("Output file created")
		defer file.Close()
		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err.Error())
	}

	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file1, err1 := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err1 == nil {
		defer file1.Close()
		encoder := json.NewEncoder(file1)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err1.Error())
	}

	file2, err2 := os.CreateTemp(".", "tempfile-*.json")
	if err2 == nil {
		defer file2.Close()
		encoder := json.NewEncoder(file2)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err2.Error())
	}

	path, err3 := os.UserHomeDir()
	if err3 == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	err4 := os.MkdirAll(filepath.Dir(path), 0766)
	if err4 == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err4 != nil {
		Printfln("Error %v", err4.Error())
	}

	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))

	path1, err5 := os.Getwd()
	if err5 == nil {
		dirEntries, err := os.ReadDir(path1)
		if err == nil {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}
	if err5 != nil {
		Printfln("Error %v", err5.Error())
	}

	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}
	}

	path2, err6 := os.Getwd()
	if err6 == nil {
		matches, err := filepath.Glob(filepath.Join(path2, "*.json"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err6 != nil {
		Printfln("Error %v", err6.Error())
	}

	path3, err7 := os.Getwd()
	if err7 == nil {
		err = filepath.WalkDir(path3, callback)
	} else {
		Printfln("Error %v", err7.Error())
	}
}
