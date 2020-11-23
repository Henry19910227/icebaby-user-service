package model

// User ...
type User struct {
	ID           int64
	Role         int
	InviteCode   string
	InviteUserID int64
	Status       int
	CreateAt     string
	UpdateAt     string
	Nickname     string
	Avatar       string
	Intro        string
	Sex          int
	Birthday     string
	Email        string
	Area         string
	Height       int
	Weight       int
	Favorite     string
	Smoke        int
	Drink        int
	AuthType     int
	Identifier   string
	Password     string
}
