package model

import (
	"time"

	"github.com/gocql/gocql"
)

type Order struct {
	ID         gocql.UUID `json:"order_id"`
	CustomerID gocql.UUID `json:"customer_id"`
	OrderDate  time.Time  `json:"order_date"`
	Amount     float64    `json:"total_amount"`
}
