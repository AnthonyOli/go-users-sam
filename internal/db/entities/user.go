package entities

import "time"

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string `db:"password_hash"`
	Phone        string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
