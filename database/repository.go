package database

import (
	"fmt"

	"github.com/RaphaSalomao/alura-webapp-go/model"
	"github.com/google/uuid"
)

func FindAllProducts() []model.Product {
	db := Connect()
	defer db.Close()
	rawProducts, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic(err)
	}

	products := []model.Product{}

	for rawProducts.Next() {
		var id uuid.UUID
		var name, description string
		var price float64

		var quantity int

		err = rawProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err)
		}
		products = append(products, model.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		})
	}

	return products
}

func FindProduct(id uuid.UUID) model.Product {
	db := Connect()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		panic(err)
	}
	var productId uuid.UUID
	var name, description string
	var price float64
	var quantity int

	err = stmt.QueryRow(id).Scan(&productId, &name, &description, &price, &quantity)
	if err != nil {
		panic(err.Error())
	}
	return model.Product{
		Id:          productId,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}

func CreateProduct(p *model.Product) {
	db := Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO product (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.Quantity)
	if err != nil {
		panic(err)
	}
}

func DeleteProduct(id uuid.UUID) {
	db := Connect()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func UpdateProduct(p *model.Product) {
	fmt.Println("Updating product...", p)
	db := Connect()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE product SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.Quantity, p.Id)
	if err != nil {
		panic(err)
	}
}
