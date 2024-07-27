-- Insert data into products table
INSERT INTO products (name, price, stock) VALUES
('Product A', 10000.00, 50),
('Product B', 15000.00, 30),
('Product C', 20000.00, 20);

-- Insert data into staff table
INSERT INTO staff (name, email, position) VALUES
('Kamal', 'kamal@espresso.com', 'Manager'),
('Bagas', 'bagas@espresso.com', 'Sales'),
('Marcel', 'marcel@espresso.com', 'Support');

-- Insert data into sales table
INSERT INTO sales (product_id, quantity, sale_date) VALUES
(1, 5, '2024-07-01'),
(2, 3, '2024-07-02'),
(3, 2, '2024-07-03');
