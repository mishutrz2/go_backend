package models

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// create new user
func (m *DBModel) CreateAccount(usr User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 12)
	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt := `insert into users (nickname, full_name, country, email, password) values ($1, $2, $3, $4, $5)`
	_, err = m.DB.ExecContext(ctx, stmt, usr.Nickname, usr.FullName, usr.Country, usr.Email, hashedPassword)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// get user info
func (m *DBModel) GetUserInfo(emailAdress string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	query := `select id_user, nickname, full_name, country, email, password from users where users.email=$1`
	row := m.DB.QueryRowContext(ctx, query, emailAdress)

	err := row.Scan(
		&user.ID_user,
		&user.Nickname,
		&user.FullName,
		&user.Country,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

// get user info from id
func (m *DBModel) GetUserPublicInfo(id_user int) (*UserNoPass, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user UserNoPass
	query := `select nickname, full_name, country, email from users where users.id_user=$1`
	row := m.DB.QueryRowContext(ctx, query, id_user)

	err := row.Scan(
		&user.Nickname,
		&user.FullName,
		&user.Country,
		&user.Email,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
