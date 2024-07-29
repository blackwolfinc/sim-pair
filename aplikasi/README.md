# Inventory Management System

This is a simple CLI-based inventory management system built with Go, using MySQL as the database.

## Features

- **Add Product:** Allows users to add new products to the database.
- **Update Product Stock:** Allows users to update the stock of existing products.
- **Add Staff:** Allows users to add new staff members to the database.
- **Sales Report:** Displays sales reports based on a specified date range.
- **Exit:** Exits the application.

## Installation

### Prerequisites

- Go installed on your machine. [Download Go](https://golang.org/dl/)
- MySQL installed on your machine. [Download MySQL](https://dev.mysql.com/downloads/mysql/)

### Steps

1. **Clone the repository:**

    ```sh
    git clone https://github.com/blackwolfinc/sim-pair
    cd sim-pair
    ```

2. **Create `.env` file:**

    Create a `.env` file in the root directory of your project with the following content:

    ```env
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=
    DB_NAME=SimPair
    ```

    Replace the placeholder values with your actual MySQL database credentials.

3. **Install dependencies:**

    ```sh
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/joho/godotenv
    ```

4. **Set up the database:**

    Create the database and tables using the provided DDL statements in `ddl.sql`.

    ```sql
    -- Create table for products
    CREATE TABLE products (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        stock INT NOT NULL
    );

    -- Create table for staff
    CREATE TABLE staff (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        position VARCHAR(255) NOT NULL
    );

    -- Create table for sales
    CREATE TABLE sales (
        id INT AUTO_INCREMENT PRIMARY KEY,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        sale_date DATE NOT NULL,
        FOREIGN KEY (product_id) REFERENCES products(id)
    );
    ```

5. **Run the application:**

    ```sh
    go run main.go
    ```

## Database Schema

### Products Table

- `id` (auto-incremented): Primary key.
- `name`: Product name.
- `price`: Product price.
- `stock`: Available stock.

### Staff Table

- `id` (auto-incremented): Primary key.
- `name`: Staff name.
- `email`: Staff email.
- `position`: Staff position.

### Sales Table

- `id` (auto-incremented): Primary key.
- `product_id`: Foreign key referencing `products(id)`.
- `quantity`: Quantity sold.
- `sale_date`: Date of sale.

## Data Definition Language (DDL)

```sql
-- Create table for products
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL
);

-- Create table for staff
CREATE TABLE staff (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL
);

-- Create table for sales
CREATE TABLE sales (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    sale_date DATE NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);
```

## Pembagian Tugas
Marcel: fitur 1 dan 2
Kamal: fitur 3
Bagas: fitur 4
Untuk role, kami menjalankannya secara bergilir di mana semua orang melakukan semua role yang ada, dengan pengaplikasian metode pull request di github.
Notes: Untuk yang lainnya kita kerjakan bersama
