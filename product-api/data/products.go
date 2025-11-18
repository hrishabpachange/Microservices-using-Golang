package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json: "id"`
	Name        string  `json: "name`
	Description string  `json: "description`
	Price       float32 `json: "price`
	SKU         string  `json: "sku"`
	CreatedOn   string  `json: "-"`
	UpdatedOn   string  `json: "-"`
	DeletedOn   string  `json: "-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p) 
}

func GetProducts() []*Product {
	return ProductList
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       60,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Coffee without milk",
		Price:       40,
		SKU:         "dfgsdkg",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
