package model

import "database/sql"

// User ...
type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Userinfo *Userinfo `json:"userinfo"`
}

// NewUser ...
func NewUser(uid int64, email string, password string, name sql.NullString, image sql.NullString, birthday sql.NullString) *User {

	if name.Valid {
		return &User{
			ID:       uid,
			Email:    email,
			Password: password,
			Userinfo: &Userinfo{
				Name:     name.String,
				Birthday: birthday.String,
				Image:    image.String,
			},
		}
	}
	return &User{
		ID:       uid,
		Email:    email,
		Password: password,
	}
}
