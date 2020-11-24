package model

// APILoginRequest ...
type APILoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
}
