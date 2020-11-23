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

// Validate 驗證帳號密碼是否存在
func (repo *ICValidateRepository) Validate(identifier string, password string) (int64, error) {
	query := "SELECT users.id FROM users\n" +
		"INNER JOIN user_auths\n" +
		"ON users.id = user_auths.user_id\n" +
		"WHERE user_auths.identifier = ? AND user_auths.`password` = ?"
	row := repo.db.QueryRow(query, identifier, password)
	var uid sql.NullInt64
	err := row.Scan(&uid)
	if err != nil {
		return 0, err
	}
	if !uid.Valid {
		return 0, errors.New("帳號或密碼錯誤")
	}
	return uid.Int64, nil
}
