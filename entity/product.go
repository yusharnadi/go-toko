package entity

import "time"

type Product struct {
	ID        uint64
	Name      string
	Price     int64
	Stock     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
