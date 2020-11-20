package model

// User ...
type User struct {
	ID           int64  `json:"id"`
	Role         string `json:"role"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Intro        string `json:"intro"`
	Sex          int    `json:"sex"`
	InviteCode   string `json:"invite_code"`
	InviteUserID string `json:"invite_user_id"`
	AuthType     int    `json:"auth_type"`
	Identifier   string `json:"identifier"`
	Password     string `json:"password"`
}
