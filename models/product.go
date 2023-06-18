package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Product_id                uuid.UUID
	Productname               string          `json:"product_name"`
	Productdescription        string          `json:"product_description"`
	Productimages             []string        `json:"product_images"`
	Productprice              decimal.Decimal `json:"product_price"`
	Compressed_product_images []string
	Createdat                 time.Time
	Updatedat                 time.Time
}
