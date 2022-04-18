package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

// var validUser = models.User{
// 	ID_user:  18,
// 	Nickname: "shumi",
// 	FullName: "Mircea",
// 	Country:  "Romania",
// 	Email:    "hello@yahoo.com",
// 	Password: "$2a$12$squqQj9ISEdVV9c32LbSQeU/0XloPvd2LuUWHeQnCoLjSrfvxIjme",
// }

// to verify login
type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type LoggedInInfo struct {
	Token    string `json:"token"`
	Nickname string `json:"nickname"`
	Id_user  int    `json:"id_user"`
	Country  string `json:"country"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// var result LoggedInInfo

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		app.errJSON(w, errors.New("unauthorized"))
		return
	}

	user, err := app.models.DB.GetUserInfo(creds.Username)

	if err != nil {
		app.errJSON(w, errors.New("problem with credentials"))
	}
	// query in database for hashed password
	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errJSON(w, errors.New("unauthorized"))
		return
	}

	// we have valid user and password

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(user.ID_user)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "ceva.com"
	claims.Audiences = []string{"ceva.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errJSON(w, errors.New("error signing"))
		return
	}

	// result.Token = string(jwtBytes)
	// result.Id_user = user.ID_user
	// result.Nickname = user.Nickname
	// result.Country = user.Country

	app.writeJSON(w, http.StatusOK, string(jwtBytes), "response")

}
