package models

import "github.com/shopspring/decimal"

type Product struct {
	Productname        string          `json:"product_name"`
	Productdescription string          `json:"product_description"`
	Productimages      []string        `json:"product_images"`
	Productprice       decimal.Decimal `json:"product_price"`
}
