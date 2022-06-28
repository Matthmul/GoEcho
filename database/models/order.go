package models

type Order struct {
	ID 			int
	OrderNumber int
	CartID 		int
	Cart     	Cart
	PaymentID 	int
	Payment     Payment
}