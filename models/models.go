package models

import (
	"database/sql"
	"time"
)

// wrapper for database
type Models struct {
	DB DBModel
}

// return models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// this is the type for prediction
type Prediction struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      int       `json:"author"`
	Votes       int       `json:"votes"`
	Winner      string    `json:"winner"`
}

type DisplayPrediction struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      string    `json:"author"`
	Votes       int       `json:"votes"`
	Winner      string    `json:"winner"`
}

type User struct {
	ID_user  int    `json:"id_user"`
	Nickname string `json:"nickname"`
	FullName string `json:"full_name"`
	Country  string `json:"country"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserNoPass struct {
	Nickname string `json:"nickname"`
	FullName string `json:"full_name"`
	Country  string `json:"country"`
	Email    string `json:"email"`
}
