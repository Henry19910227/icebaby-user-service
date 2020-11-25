package model

// APIUserDetailRequest GetUserDetail 請求
type APIUserDetailRequest struct {
	ID int64 `json:"id" binding:"required"`
}

// UserAll 全部用戶資訊
type UserAll struct {
	ID           int64
	Role         int
	Status       int
	CreateAt     string
	UpdateAt     string
	LastLogin    string
	LoginStatus  int
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
	InviteCode   string
	InviteUserID int64
	AuthType     int
	Identifier   string
	Password     string
}

// User 用戶基本資訊
type User struct {
	ID       int64
	Nickname string
	Avatar   string
	Intro    string
	Birthday string
}

// NewUser ...
func NewUser(userAll *UserAll) *User {
	return &User{
		ID:       userAll.ID,
		Nickname: userAll.Nickname,
		Avatar:   userAll.Avatar,
		Intro:    userAll.Intro,
		Birthday: userAll.Birthday,
	}
}

// UserDetail 用戶詳細
type UserDetail struct {
	ID       int64
	Nickname string
	Avatar   string
	Intro    string
	Sex      int
	Birthday string
	Email    string
	Area     string
	Height   int
	Weight   int
	Favorite string
	Smoke    int
	Drink    int
}

// NewUserDetail 創建 UserDetail
func NewUserDetail(userAll *UserAll) *UserDetail {
	return &UserDetail{
		ID:       userAll.ID,
		Nickname: userAll.Nickname,
		Avatar:   userAll.Avatar,
		Intro:    userAll.Intro,
		Sex:      userAll.Sex,
		Birthday: userAll.Birthday,
		Email:    userAll.Email,
		Area:     userAll.Area,
		Height:   userAll.Height,
		Weight:   userAll.Weight,
		Favorite: userAll.Favorite,
		Smoke:    userAll.Smoke,
		Drink:    userAll.Drink,
	}
}
