# SIMPAIR

This is a simple CLI-based system built with Go, using MySQL as the database.

## Features

- Add Product: Allows users to add new products to the database.
- Update Product Stock: Allows users to update the stock of existing products.
- Add Staff: Allows users to add new staff members to the database.
- Sales Report: Displays sales reports based on a specified date range.
- Exit: Exits the application.

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
    DB_NAME=simpair
    ```

    Replace the placeholder values with your actual MySQL database credentials.

3. **Install dependencies:**

    ```sh
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/joho/godotenv
    ```

4. **Set up the database:**

    Create the database and tables using the provided DDL and DML statements.

5. **Run the application:**

    ```sh
    cd aplikasi\cli
    go run main.go
    ```
