package repository

import (
	"database/sql"
	"errors"

	"github.com/Henry19910227/icebaby-user-service/internal/model"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository 創建一個 UserRepository
func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{conn}
}

// GetAll Implement UserRepository interface
func (ur *userRepository) GetAll() ([]*model.User, error) {
	query := "SELECT users.id,users.email,users.password,userinfo.name,userinfo.image,userinfo.birthday\n" +
		"FROM users\n" +
		"LEFT JOIN userinfo\n" +
		"ON users.userinfo_id = userinfo.id "
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	users := []*model.User{}
	for rows.Next() {
		var uid int64
		var email string
		var password string
		var name sql.NullString
		var image sql.NullString
		var birthday sql.NullString
		if err := rows.Scan(&uid, &email, &password, &name, &image, &birthday); err == nil {
			user := model.NewUser(uid, email, password, name, image, birthday)
			users = append(users, user)
		}
	}
	return users, nil
}

// GetUser 以 email 與 password 獲取用戶
func (ur *userRepository) GetUser(email string, password string) (*model.User, error) {
	query := "SELECT id FROM users WHERE email = ? AND password = ?"
	row := ur.db.QueryRow(query, email, password)
	var uid int64
	if err := row.Scan(&uid); err != nil {
		return nil, err
	}
	return ur.GetByID(uid)

}

// GetById ...
func (ur *userRepository) GetByID(id int64) (*model.User, error) {
	query := "SELECT users.id,users.email,users.password,userinfo.name,userinfo.image,userinfo.birthday\n" +
		"FROM users\n" +
		"LEFT JOIN userinfo\n" +
		"ON users.userinfo_id = userinfo.id\n" +
		"WHERE users.id = ?"
	row := ur.db.QueryRow(query, id)

	var uid int64
	var email string
	var password string
	var name sql.NullString
	var image sql.NullString
	var birthday sql.NullString
	if err := row.Scan(&uid, &email, &password, &name, &image, &birthday); err != nil {
		return nil, err
	}
	return model.NewUser(uid, email, password, name, image, birthday), nil
}

// Add ...
func (ur *userRepository) Add(email string, password string, name string, birthday string) (int64, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	query := "INSERT INTO userinfo (name,image,birthday) VALUES (?,?,?)"
	infoRes, err := tx.Exec(query, name, "", birthday)
	if err != nil {
		return 0, err
	}
	infoLastID, err := infoRes.LastInsertId()
	if err != nil {
		return 0, err
	}
	query = "INSERT INTO users (email,password,userinfo_id) VALUES (?,?,?)"
	userRes, err := tx.Exec(query, email, password, infoLastID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return userRes.LastInsertId()
}

// DeleteByID ...
func (ur *userRepository) DeleteByID(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	res, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("查無此用戶")
	}
	return nil
}

// UpdateUserinfo ...
func (ur *userRepository) UpdateUserinfo(uid int64, name string, birthday string) (*model.User, error) {
	query := "UPDATE userinfo\n" +
		"INNER JOIN users ON userinfo.id = users.userinfo_id\n" +
		"SET userinfo.name = ?,userinfo.birthday = ?\n" +
		"WHERE users.id = ?"
	_, err := ur.db.Exec(query, name, birthday, uid)
	if err != nil {
		return nil, err
	}
	return ur.GetByID(uid)
}

// UpdateEmail ...
func (ur *userRepository) UpdateEmail(uid int64, email string) (*model.User, error) {
	query := "UPDATE users SET email = ? WHERE id = ?"
	if _, err := ur.db.Exec(query, email, uid); err != nil {
		return nil, err
	}
	return ur.GetByID(uid)
}

// UpdatePassword ...
func (ur *userRepository) UpdatePassword(uid int64, password string) error {
	query := "UPDATE users SET password = ? WHERE id = ?"
	_, err := ur.db.Exec(query, password, uid)
	return err
}

// UpdateUserImage ...
func (ur *userRepository) UpdateUserImage(uid int64, image string) (*model.User, error) {
	query := "UPDATE userinfo\n" +
		"INNER JOIN users ON userinfo.id = users.userinfo_id\n" +
		"SET userinfo.image = ?\n" +
		"WHERE users.id = ?"
	if _, err := ur.db.Exec(query, image, uid); err != nil {
		return nil, err
	}
	return ur.GetByID(uid)
}
