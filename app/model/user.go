package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRegister struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}

func (model *Model) AddNewUser(user UserRegister) (string, error) {
	args := pgx.NamedArgs{
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"name":     user.Name,
		"phone":    user.PhoneNumber,
	}

	query := `
		INSERT INTO splitto.user (username, password, email, name, phone)
		VALUES (@username, @password, @email, @name, @phone)
		RETURNING *
	`

	conn := model.DBConn
	_, err := conn.Exec(context.Background(), query, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			switch pgErr.Code {
			case pgerrcode.UniqueViolation:

			}
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}

		fmt.Println("Add New User failed: %v", err)
	}

	return user_id, err
}
