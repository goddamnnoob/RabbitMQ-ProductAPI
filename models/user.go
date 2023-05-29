package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	Userid    int32 `json:"user_id"`
	Name      string
	Mobile    string
	Latitude  decimal.Decimal
	Longitude decimal.Decimal
	Createdat time.Time
	Updatedat time.Time
}
