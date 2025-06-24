package models

import (
	"context"
	"fmt"
	"minitask2/utils"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id int `json:"id"`

	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  string `json:"errors,omiempty"`
	Result  any    `json:"result,omiempty"`
}

var Data []Profile
var OtpStore = make(map[string]string)

func (p Profile) Register() {

}
func FindAllUser(search string) []Profile {

	conn, err := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	if err != nil {

	}
	rows, err := conn.Query(
		context.Background(), `SELECT * FROM users WHERE username ILIKE $1`, fmt.Sprintf("%%%s%%", search))
	users, err := pgx.CollectRows[Profile](rows, pgx.RowToStructByName)
	if err != nil {

	}
	return users
}
func FindUserById(id string) Profile {
	conn, err := utils.DBConnect()
	if err != nil {

	}
	rows, err := conn.Query(context.Background(), `
	SELECT * FROM users WHERE id = $1
	`, id)
	users, err := pgx.CollectOneRow[Profile](rows, pgx.RowToStructByName)
	return users
}
func CreateUser(user Profile) (Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Profile{}, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	_, err = conn.Exec(context.Background(), `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
	`, user.Username, user.Email, user.Password)

	if err != nil {
		return Profile{}, err
	}

	return user, nil
}
func UpdateUser(id string, newData Profile) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	oldUser := FindUserById(id)

	if newData.Username != "" {
		oldUser.Username = newData.Username
	}
	if newData.Email != "" {
		oldUser.Email = newData.Email
	}
	if newData.Password != "" {
		oldUser.Password = newData.Password
	}

	_, err = conn.Exec(context.Background(), `
		UPDATE users SET username = $1, email = $2, password = $3
		WHERE id = $4
	`, oldUser.Username, oldUser.Email, oldUser.Password, id)

	return err
}
func DeleteUser(id string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	getUser := FindUserById(id)
	if getUser.Id == 0 {
		return fmt.Errorf("user dengan ID %s tidak ditemukan", id)
	}

	_, err = conn.Query(context.Background(), `
		DELETE FROM users WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}
