package model

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Price       float64
	Quantity    int
}
