package main

import (
	"aplikasi/config"
	"aplikasi/fitur"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up database connection
	db, err := sql.Open("mysql", config.DatabaseConfig())
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
		fmt.Scanln(&choice)

		reader := bufio.NewReader(os.Stdin)
		switch choice {
		case 1:
			//addProduct(db)
			fmt.Print("Enter a product name: ")
			productName, _ := reader.ReadString('\n')
			productName = strings.TrimSpace(productName)

			fmt.Print("Enter product price: ")
			priceInput, _ := reader.ReadString('\n')
			priceInput = strings.TrimSpace(priceInput)
			price, err := strconv.ParseFloat(priceInput, 64)
			if err != nil {
				fmt.Println("Invalid price input:", err)
				return
			}

			fmt.Print("Enter product stock: ")
			stockInput, _ := reader.ReadString('\n')
			stockInput = strings.TrimSpace(stockInput)
			stock, err := strconv.Atoi(stockInput)
			if err != nil {
				fmt.Println("Invalid stock input:", err)
				return
			}
			fitur.AddProduct(db, productName, price, stock)
		case 2:
			//updateProductStock(db)
			fmt.Print("Enter the product name you want to update: ")
			productName, _ := reader.ReadString('\n')
			productName = strings.TrimSpace(productName)

			fmt.Print("Do you want to add or decrease the stock (a/d): ")
			addStockInp, _ := reader.ReadString('\n')
			addStockInp = strings.TrimSpace(addStockInp)
			addStock := true
			if addStockInp == "d" {
				addStock = false
			}

			fmt.Print("Enter product stock: ")
			stockInput, _ := reader.ReadString('\n')
			stockInput = strings.TrimSpace(stockInput)
			stock, err := strconv.Atoi(stockInput)
			if err != nil {
				fmt.Println("Invalid stock input:", err)
				return
			}
			fitur.UpdateProductStock(db, productName, stock, addStock)
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
