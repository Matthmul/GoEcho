package models

type Cart struct {
	ID 			int
	CartNumber	int
	ProductID 	int
	Product     Product
	Quantity    int
}