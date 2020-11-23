package model

// APILoginReq ...
type APILoginReq struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
}

// APILoginRes ...
type APILoginRes struct {
	ID       int64
	Role     int
	Status   int
	Nickname string
}
