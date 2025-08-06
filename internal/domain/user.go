package domain

import "time"

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRepository interface {
	GetById(Id string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}


