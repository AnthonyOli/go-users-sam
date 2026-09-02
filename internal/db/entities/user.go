package entities

import "time"

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string `db:"password_hash"`
	Phone        string
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
