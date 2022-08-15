package domain

import "time"

type User struct {
	ID        int64
	Nickname  string `gorm:"size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
