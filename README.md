## GoEcho

# Run code
```
go run server.go 
```

## Endpoints

- Cart
```
	ID 			int
	CartNumber	int
	ProductID 	int
	Product     Product
	Quantity    int
```

- Category
```
	ID 			int
	Name 		string
```

- Order
```
	ID 			int
	OrderNumber int
	CartID 		int
	Cart     	Cart
	PaymentID 	int
	Payment     Payment
```

- Payment
```
	ID 				int
	CardNumber 		string
	PostalCode     	int
	Date      		string
	Status     		string
```

- Product
```
	ID 			int
	Name       	string
	Price      	int
	CategoryID  int
	Category   	Category
```