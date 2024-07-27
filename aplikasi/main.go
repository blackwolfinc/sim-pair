package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database connection info from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Set up database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	for {
		// Display the menu
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
			//addStaff(db)
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
