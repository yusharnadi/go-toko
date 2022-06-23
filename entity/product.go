package entity

import "time"

type Product struct {
	ID        uint64
	Name      string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
