package handler

import (
	"database/sql"
	"fmt"
	"log"

	"aplikasi/entity"
)

func AddStaff(db *sql.DB, staff entity.Staff) {
	sqlStatement := `INSERT INTO staff (name, email, position) VALUES (?, ?, ?)`
	_, err := db.Exec(sqlStatement, staff.Name, staff.Email, staff.Position)
	if err != nil {
		log.Fatalf("Failed to insert staff: %v", err)
	}
	fmt.Println("Staff berhasil ditambahkan!")
}
