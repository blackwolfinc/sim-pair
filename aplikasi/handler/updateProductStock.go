package handler

import (
	"database/sql"
	"fmt"
	"log"
)

func UpdateProductStock(db *sql.DB, name string, stock int, addStock bool) {
	var query string
	if addStock {
		query = `UPDATE products SET stock = stock + ? WHERE name = ?;`
	} else {
		query = `UPDATE products SET stock = stock - ? WHERE name = ?;`
	}

	result, err := db.Exec(query, stock, name)
	if err != nil {
		log.Fatal(err)
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if RowsAffected == 0 {
		log.Println("No rows affected, check product name")
		return
	}
	fmt.Println("=======================================================================================")
	fmt.Println("Product added Successfully")
}
