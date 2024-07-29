package main

import (
	"aplikasi/config"
	"aplikasi/fitur"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env fileß
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
		reader := bufio.NewReader(os.Stdin)

		choiceInp, _ := reader.ReadString('\n')
		choiceInp = strings.TrimSpace(choiceInp)
		choice, err := strconv.Atoi(choiceInp)
		if err != nil {
			fmt.Println("Invalid choice input:", err)
		}

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
				continue
			}

			fmt.Print("Enter product stock: ")
			stockInput, _ := reader.ReadString('\n')
			stockInput = strings.TrimSpace(stockInput)
			stock, err := strconv.Atoi(stockInput)
			if err != nil {
				fmt.Println("Invalid stock input:", err)
				continue
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
				fmt.Println("=======================================================================================")
				fmt.Println("Invalid stock input:", err)
				fmt.Println("=======================================================================================")
				continue
			}
			fitur.UpdateProductStock(db, productName, stock, addStock)
		case 3:
			//addStaff(db)
		case 4:
			var startDate, endDate string

			for {

				fmt.Print("Enter Start Date - (yyyy/mm/dd) : ")
				startDate, _ = reader.ReadString('\n')
				startDate = strings.TrimSpace(startDate)

				// Validate start date
				_, err := time.Parse("2006/01/02", startDate)
				if err != nil {
					fmt.Println("=======================================================================================")
					fmt.Println("Invalid start date format. Please use yyyy/mm/dd.")
					fmt.Println("=======================================================================================")
					continue
				}
				break
			}

			for {
				fmt.Print("Enter End Date - (yyyy/mm/dd) : ")
				endDate, _ = reader.ReadString('\n')
				endDate = strings.TrimSpace(endDate)

				// Validate end date
				_, err := time.Parse("2006/01/02", endDate)
				if err != nil {
					fmt.Println("=======================================================================================")
					fmt.Println("Invalid end date format. Please use yyyy/mm/dd.")
					fmt.Println("=======================================================================================")
					continue
				}
				break
			}

			// Call the SummarySales function
			err = fitur.SummarySales(db, startDate, endDate)
			if err != nil {
				log.Fatal(err)
			}
		case 5:
			fmt.Println("=======================================================================================")
			fmt.Println("Keluar dari aplikasi...")
			fmt.Println("=======================================================================================")
			return
		default:
			fmt.Println("=======================================================================================")
			fmt.Println("Opsi tidak valid!")
			fmt.Println("=======================================================================================")
		}
	}
}
