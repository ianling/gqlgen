// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package dataloader

import (
	time "time"
)

type Address struct {
	ID      int
	Street  string
	Country string
}
type Customer struct {
	ID        int
	Name      string
	AddressID int
}
type Item struct {
	Name string
}
type Order struct {
	ID     int
	Date   time.Time
	Amount float64
}