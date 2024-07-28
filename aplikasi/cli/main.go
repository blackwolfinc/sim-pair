package main

import (
	"database/sql"
	"fmt"
	"log"

	"aplikasi/config"
	"aplikasi/entity"
	"aplikasi/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Ubah Stok Produk")
		fmt.Println("3. Tambah Staff")
		fmt.Println("4. Rekap Penjualan")
		fmt.Println("5. Exit")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			//addProduct(db)
		case 2:
			//updateProductStock(db)
		case 3:
			addStaff(db)
		case 4:
			//reportSales(db)
		case 5:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Opsi tidak valid!")
		}
	}
}

func addStaff(db *sql.DB) {
	var name, email, position string

	fmt.Print("Nama Staff: ")
	fmt.Scan(&name)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Posisi: ")
	fmt.Scan(&position)

	staff := entity.Staff{Name: name, Email: email, Position: position}
	handler.AddStaff(db, staff)
}
