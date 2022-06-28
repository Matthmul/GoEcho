package models

type Product struct {
	ID 			int
	Name       	string
	Price      	int
	CategoryID  int
	Category   	Category
}