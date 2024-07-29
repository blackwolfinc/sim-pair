package handler

import (
	"database/sql"
	"fmt"
)

// ProductSalesDetail represents the details of each sold product
type ProductSalesDetail struct {
	ProductName string
	Quantity    int
	SaleDate    string
	TotalPrice  float64
}

// SummarySales generates a detailed sales summary report for a given date range
func SummarySales(db *sql.DB, startDate string, endDate string) error {
	// Execute the query
	query := `
	SELECT 
		p.name AS product_name,
		s.quantity,
		s.sale_date,
		s.quantity * p.price AS total_price
	FROM 
		sales s
	JOIN 
		products p ON s.product_id = p.id
	WHERE 
		s.sale_date BETWEEN ? AND ?
	ORDER BY
		s.sale_date;
	`

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return err
	}
	defer rows.Close()

	var details []ProductSalesDetail
	var totalProductsSold, totalSalesTransactions int
	var totalRevenue float64

	for rows.Next() {
		var detail ProductSalesDetail
		err := rows.Scan(&detail.ProductName, &detail.Quantity, &detail.SaleDate, &detail.TotalPrice)
		if err != nil {
			return err
		}
		details = append(details, detail)
		totalProductsSold += detail.Quantity
		totalSalesTransactions++
		totalRevenue += detail.TotalPrice
	}

	if err = rows.Err(); err != nil {
		return err
	}

	// Output the detailed results
	fmt.Println("=======================================================================================")
	fmt.Println("Sales Summary:")
	fmt.Println("=======================================================================================")

	for _, detail := range details {
		fmt.Printf("Product: %s, Quantity: %d, Sale Date: %s, Total Price: %.2f\n", detail.ProductName, detail.Quantity, detail.SaleDate, detail.TotalPrice)
	}
	fmt.Println("=======================================================================================")

	// Output the aggregate results
	fmt.Printf("\nTotal Products Sold: %d\n", totalProductsSold)
	fmt.Printf("Total Sales Transactions: %d\n", totalSalesTransactions)
	fmt.Printf("Total Revenue: %.2f\n", totalRevenue)

	fmt.Println("=======================================================================================")

	return nil
}
