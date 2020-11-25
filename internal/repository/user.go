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

// Add ...
func (ur *userRepository) InsertUser(user *model.UserAll) (int64, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	query := "INSERT INTO users (role,status) VALUES (?,?)"
	userRes, err := tx.Exec(query, user.Role, user.Status)
	if err != nil {
		return 0, err
	}
	lastUserID, err := userRes.LastInsertId()
	if err != nil {
		return 0, err
	}
	query = "INSERT INTO user_auths (type,identifier,password,user_id) VALUES (?,?,?,?)"
	_, err = tx.Exec(query, user.AuthType, user.Identifier, user.Password, lastUserID)
	if err != nil {
		return 0, err
	}
	query = "INSERT INTO user_details (user_id,nickname,avatar,intro,sex,birthday,email,area,height,weight,favorite,smoke,drink,invite_code,invite_user_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err = tx.Exec(query, lastUserID, user.Nickname, user.Avatar, user.Intro, user.Sex, user.Birthday, user.Email, user.Area, user.Height, user.Weight, user.Favorite, user.Smoke, user.Drink, user.InviteCode, user.InviteUserID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return userRes.LastInsertId()
}

// GetUserIDByCode 以 inviteCode 獲取 uid
func (ur *userRepository) GetUserIDByCode(inviteCode string) (int64, error) {
	query := "SELECT id FROM users WHERE invite_code = ?"
	row := ur.db.QueryRow(query, inviteCode)
	var uid int64
	if err := row.Scan(&uid); err != nil {
		return 0, errors.New("無效的邀請碼")
	}
	return uid, nil
}

// GetAll Implement UserRepository interface
func (ur *userRepository) GetAll() ([]*model.UserAll, error) {
	query := "SELECT users.id, users.role, users.invite_code, users.invite_user_id, users.`status`, users.create_at, users.update_at,\n" +
		"user_details.nickname, user_details.avatar, user_details.intro, user_details.sex,\n" +
		"user_details.birthday, user_details.email, user_details.area, user_details.height,\n" +
		"user_details.weight, user_details.favorite, user_details.smoke, user_details.drink,\n" +
		"user_auths.type, user_auths.identifier, user_auths.`password`\n" +
		"FROM users\n" +
		"INNER JOIN user_details ON users.id = user_details.user_id\n" +
		"INNER JOIN user_auths ON users.id = user_auths.user_id\n" +
		"WHERE user_auths.identifier = ? && user_auths.`password` = ?"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	users := []*model.UserAll{}
	for rows.Next() {
		var uid int64
		var email string
		var password string
		var name sql.NullString
		var image sql.NullString
		var birthday sql.NullString
		if err := rows.Scan(&uid, &email, &password, &name, &image, &birthday); err == nil {
			users = append(users, nil)
		}
	}
	return users, nil
}

// GetById ...
func (ur *userRepository) GetUserByID(id int64) (*model.UserAll, error) {
	query := "SELECT users.id, users.role, users.`status`, users.create_at, users.update_at,\n" +
		"users.login_status, users.last_login,\n" +
		"user_details.nickname, user_details.avatar, user_details.intro, user_details.sex,\n" +
		"user_details.birthday, user_details.email, user_details.area, user_details.height,\n" +
		"user_details.weight, user_details.favorite, user_details.smoke, user_details.drink,\n" +
		"user_details.invite_code, user_details.invite_user_id,\n" +
		"user_auths.type, user_auths.identifier, user_auths.`password`\n" +
		"FROM users\n" +
		"INNER JOIN user_details ON users.id = user_details.user_id\n" +
		"INNER JOIN user_auths ON users.id = user_auths.user_id\n" +
		"WHERE users.id = ?;"
	row := ur.db.QueryRow(query, id)

	var user model.UserAll
	if err := row.Scan(&user.ID, &user.Role, &user.Status, &user.CreateAt, &user.UpdateAt, &user.LoginStatus, &user.LastLogin,
		&user.Nickname, &user.Avatar, &user.Intro, &user.Sex, &user.Birthday, &user.Email, &user.Area, &user.Height, &user.Weight,
		&user.Favorite, &user.Smoke, &user.Drink, &user.InviteCode, &user.InviteUserID, &user.AuthType, &user.Identifier, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
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
func (ur *userRepository) UpdateUserinfo(uid int64, name string, birthday string) (*model.UserAll, error) {
	query := "UPDATE userinfo\n" +
		"INNER JOIN users ON userinfo.id = users.userinfo_id\n" +
		"SET userinfo.name = ?,userinfo.birthday = ?\n" +
		"WHERE users.id = ?"
	_, err := ur.db.Exec(query, name, birthday, uid)
	if err != nil {
		return nil, err
	}
	return ur.GetUserByID(uid)
}

// UpdateEmail ...
func (ur *userRepository) UpdateEmail(uid int64, email string) (*model.UserAll, error) {
	query := "UPDATE users SET email = ? WHERE id = ?"
	if _, err := ur.db.Exec(query, email, uid); err != nil {
		return nil, err
	}
	return ur.GetUserByID(uid)
}

// UpdatePassword ...
func (ur *userRepository) UpdatePassword(uid int64, password string) error {
	query := "UPDATE users SET password = ? WHERE id = ?"
	_, err := ur.db.Exec(query, password, uid)
	return err
}

// UpdateUserImage ...
func (ur *userRepository) UpdateUserImage(uid int64, image string) (*model.UserAll, error) {
	query := "UPDATE userinfo\n" +
		"INNER JOIN users ON userinfo.id = users.userinfo_id\n" +
		"SET userinfo.image = ?\n" +
		"WHERE users.id = ?"
	if _, err := ur.db.Exec(query, image, uid); err != nil {
		return nil, err
	}
	return ur.GetUserByID(uid)
}
