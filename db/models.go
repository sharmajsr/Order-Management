// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type Order struct {
	ID              int64  `json:"id"`
	ProductID       int64  `json:"product_id"`
	UserID          int64  `json:"user_id"`
	Price           int64  `json:"price"`
	Status          string `json:"status"`
	QuantityOrdered int64  `json:"quantity_ordered"`
}

type Product struct {
	ID                int64  `json:"id"`
	Price             int64  `json:"price"`
	Details           string `json:"details"`
	QuantityAvailable int32  `json:"quantity_available"`
}

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
