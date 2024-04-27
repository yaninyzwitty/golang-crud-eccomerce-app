package model

import "github.com/gocql/gocql"

type Review struct {
	ID         gocql.UUID `json:"review_id"`
	Comment    string     `json:"comment"`
	CustomerID gocql.UUID `json:"customer_id"`
	ProductID  gocql.UUID `json:"product_id"`
	Rating     int        `json:"rating"`
}
