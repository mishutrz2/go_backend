package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
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

// return one Prediction, and error, if any
func (m *DBModel) GetSec(id int) (*Prediction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var prediction Prediction
	query := `select id, title, description, author , votes from users, predis where users.id_user=predis.author and id=$1`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&prediction.ID,
		&prediction.Title,
		&prediction.Description,
		&prediction.Author,
		&prediction.Votes,
	)

	if err != nil {
		return nil, err
	}
	return &prediction, nil
}

// return one Prediction, and error, if any
func (m *DBModel) Get(id int) (*DisplayPrediction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var prediction DisplayPrediction
	query := `select id, title, description, users.nickname , votes, created_at, updated_at from users, predis where users.id_user=predis.author and id=$1`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&prediction.ID,
		&prediction.Title,
		&prediction.Description,
		&prediction.Author,
		&prediction.Votes,
		&prediction.CreatedAt,
		&prediction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &prediction, nil
}

// return all Predictions, and error, if any
func (m *DBModel) All() ([]*Prediction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, author, votes, created_at, updated_at from predis, users where predis.author=users.id_user;`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var predictions []*Prediction

	for rows.Next() {
		var prediction Prediction
		err := rows.Scan(
			&prediction.ID,
			&prediction.Title,
			&prediction.Description,
			&prediction.Author,
			&prediction.Votes,
			&prediction.CreatedAt,
			&prediction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		predictions = append(predictions, &prediction)

	}
	return predictions, nil
}

func (m *DBModel) DbUserIdOf(id_pred int) int {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select author from predis where id=$1`

	rows, err := m.DB.QueryContext(ctx, query, id_pred)
	if err != nil {
		return 0
	}
	defer rows.Close()
	var predId int
	for rows.Next() {

		err := rows.Scan(
			&predId,
		)
		if err != nil {
			return 0
		}
	}

	return predId
}

// return all MY Predictions, and error, if any
func (m *DBModel) MyAll(id_user int) ([]*Prediction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, author, votes, created_at, updated_at from predis, users where predis.author=users.id_user and users.id_user=$1`

	rows, err := m.DB.QueryContext(ctx, query, id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var predictions []*Prediction

	for rows.Next() {
		var prediction Prediction
		err := rows.Scan(
			&prediction.ID,
			&prediction.Title,
			&prediction.Description,
			&prediction.Author,
			&prediction.Votes,
			&prediction.CreatedAt,
			&prediction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		predictions = append(predictions, &prediction)

	}
	return predictions, nil
}

func (m *DBModel) InsertPrediction(prediction Prediction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into predis (title, description, created_at, updated_at, author, votes) values ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.ExecContext(ctx, stmt, prediction.Title, prediction.Description, prediction.CreatedAt, prediction.UpdatedAt, prediction.Author, prediction.Votes)

	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) EditPrediction(prediction Prediction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update predis set title=$1, description=$2,  updated_at=$3 where id=$4`

	_, err := m.DB.ExecContext(ctx, stmt, prediction.Title, prediction.Description, prediction.UpdatedAt, prediction.ID)

	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) DeletePrediction(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from predis where id=$1`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
