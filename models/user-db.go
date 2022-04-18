package models

import (
	"context"
	"fmt"
	"time"
)

// get user info
func (m *DBModel) CreateAccount(usr User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// hash password

	stmt := `insert into users (nickname, full_name, country, email, password) values ($1, $2, $3, $4, $5)`
	_, err := m.DB.ExecContext(ctx, stmt, usr.Nickname, usr.FullName, usr.Country, usr.Email, usr.Password)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
