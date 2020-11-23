package model

// LoginInput ...
type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
}

// LoginOutput ...
type LoginOutput struct {
	ID       int64
	Role     int
	Status   int
	Nickname string
}
