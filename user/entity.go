package user

import "time"

type User struct {
	ID             int
	Occupation     string
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}