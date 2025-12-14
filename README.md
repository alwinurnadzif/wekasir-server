# Wekasir

## Installation

1. clone this repo

2. install dependency

```bash
go mod tidy
```

3. run the product

```bash
go run main.go
```

## Database structure

users / cashier

```
id: number
name: string
email: string
password: string
created_at: timestamp
updated_at: timestamp
```

products

```
id: number
name: string
description: string
qty: number
price: number
created_at: timestamp
updated_at: timestamp
```

customers

```
id: number
name: string
description: string
created_at: timestamp
updated_at: timestamp
```

transactions

```
id: number
user_id: number
customer_id: number
date: date
total_amt: number
total_qty: number
created_at: timestamp
updated_at: timestamp
```

transaction_details

```
id: number
transaction_id: number
product_id: number
qty: number
price: number
created_at: timestamp
updated_at: timestamp
```
