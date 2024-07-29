package handler

import (
	"database/sql"
	"fmt"
	"log"
)

func AddProduct(db *sql.DB, name string, price float64, stock int) {
	query := `INSERT INTO products (name, price, stock) VALUES (?, ?, ?)`
	_, err := db.Exec(query, name, price, stock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=======================================================================================")
	fmt.Println("Product added Successfully")
}
