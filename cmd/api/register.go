package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// type UserPayload struct {
// 	FullName string `json:"fullname"`
// 	Nickname string `json:"nickname"`
// 	Country  string `json:"country"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func (app *application) registerAccount(w http.ResponseWriter, r *http.Request) {
	var payload models.User

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// error handling
		app.errJSON(w, err)
	}

	// call database method
	err = app.models.DB.CreateAccount(payload)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.errJSON(w, err)
		return
	}
}

// get some info about user
func (app *application) getUserPublicInfo(w http.ResponseWriter, r *http.Request) {
	middlewareParams := r.Context().Value(UserParam{}).(string)
	idUser, err := strconv.Atoi(middlewareParams)
	if err != nil {
		app.errJSON(w, errors.New("invalid user"))
		return
	}

	usr, err := app.models.DB.GetUserPublicInfo(idUser)
	if err != nil {
		app.errJSON(w, errors.New("could not retrieve user info"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, usr, "user")
	if err != nil {
		app.errJSON(w, err)
		return
	}
}
