package model

import "github.com/gocql/gocql"

type Customer struct {
	ID      gocql.UUID `json:"customer_id"`
	Address string     `json:"address"`
	Email   string     `json:"email"`
	Name    string     `json:"name"`
}
