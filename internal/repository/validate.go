package repository

import (
	"database/sql"
	"errors"
)

// ICValidateRepository ...
type ICValidateRepository struct {
	db *sql.DB
}

// NewValidateRepository ...
func NewValidateRepository(conn *sql.DB) ValidateRepository {
	return &ICValidateRepository{conn}
}

// ValidateLogin 以 identifier 和 password 驗證用戶是否存在
func (repo *ICValidateRepository) ValidateLogin(identifier string, password string) (int64, error) {
	query := "SELECT users.id FROM users\n" +
		"INNER JOIN user_auths\n" +
		"ON users.id = user_auths.user_id\n" +
		"WHERE user_auths.identifier = ? AND user_auths.`password` = ?"
	row := repo.db.QueryRow(query, identifier, password)
	var uid sql.NullInt64
	if err := row.Scan(&uid); err != nil {
		return 0, err
	}
	if !uid.Valid {
		return 0, errors.New("帳號或密碼錯誤")
	}
	return uid.Int64, nil
}

// ValidateInviteCode 以 inviteCode 驗證 user 是否存在
func (repo *ICValidateRepository) ValidateInviteCode(inviteCode string) (int64, error) {
	query := "SELECT users.id FROM users\n" +
		"INNER JOIN user_details ON users.id = user_details.user_id\n" +
		"WHERE user_details.invite_code = ?;"
	row := repo.db.QueryRow(query, inviteCode)
	var uid sql.NullInt64
	if err := row.Scan(&uid); err != nil {
		return 0, err
	}
	if !uid.Valid {
		return 0, errors.New("無效的邀請碼")
	}
	return uid.Int64, nil
}
