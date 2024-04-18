package model

import "github.com/gocql/gocql"

type Category struct {
	ID          gocql.UUID `json:"category_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}
