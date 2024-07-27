package fitur

import (
	"database/sql"
	"log"
)

func UpdateProductStock(db *sql.DB, name string, stock int, addStock bool) {
	var query string
	if addStock {
		query = `UPDATE products SET stock = stock + ? WHERE name = ?;`
	} else {
		query = `UPDATE products SET stock = stock - ? WHERE name = ?;`
	}
	
	_, err := db.Exec(query, stock, name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Product added Successfully")
}