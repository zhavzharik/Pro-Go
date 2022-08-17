package main

func main() {
	listDrivers()
	db, err := openDatabase()
	if err == nil {
		db.Close()
	} else {
		panic(err)
	}
}

// команда для создания базы данных
//./sqlite3 products.db ".read products.sql"

// команда, чтобы проверить что база данных создана и заполнена данными
//./sqlite3 products.db "select * from PRODUCTS"

// установка пакета драйверов базы данных sqlite
//go get modernc.org/sqlite
