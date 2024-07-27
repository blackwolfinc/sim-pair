package main

import (
	"database/sql"
	"log"
	"fmt"

	"aplikasi/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// koneksi
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	// Open koneksi ke db
	db, err := sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 1. Tambah Produk Memungkinkan pengguna untuk menambah produk baru ke dalam database. 
	// Pengguna akan diminta untuk memasukkan nama produk, harga, dan jumlah stok awal.
	fmt.Println("koneksi berhasil")
}
