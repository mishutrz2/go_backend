package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// GET ALL PREDICTION
func (app *application) getAllPredictions(w http.ResponseWriter, r *http.Request) {
	predictions, err := app.models.DB.All()
	if err != nil {
		app.errJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, predictions, "predictions")
	if err != nil {
		app.errJSON(w, err)
		return
	}
}

// struct for payload (get method)
type PredictionPayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Author      string `json:"author"`
	Votes       string `json:"votes"`
}

// // CREATE
// func (app *application) createPrediction(w http.ResponseWriter, r *http.Request) {
// 	var payload PredictionPayload

// 	middlewareParams := r.Context().Value(UserParam{}).(string)
// 	idUser, err := strconv.Atoi(middlewareParams)
// 	if err != nil {
// 		app.errJSON(w, errors.New("invalid user"))
// 		return
// 	}

// 	// this will be added to database
// 	var prediction models.Prediction

// 	// prediction.ID, _ = strconv.Atoi(payload.ID)
// 	prediction.Title = payload.Title
// 	prediction.Description = payload.Description
// 	prediction.CreatedAt = time.Now()
// 	prediction.UpdatedAt = time.Now()
// 	prediction.Author = idUser

// 	err = app.models.DB.InsertPrediction(prediction)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"predis_un\"") {
// 			app.errJSON(w, errors.New("title already exists, choose another one"))
// 		} else {
// 			app.errJSON(w, err)
// 		}
// 		return
// 	}

// 	ok := jsonResp{
// 		OK: true,
// 	}

// 	err = app.writeJSON(w, http.StatusOK, ok, "response")

// 	if err != nil {
// 		app.errJSON(w, err)
// 		return
// 	}
// }

// UPDATE
func (app *application) editPrediction(w http.ResponseWriter, r *http.Request) {
	var payload PredictionPayload

	middlewareParams := r.Context().Value(UserParam{}).(string)
	idUser, err := strconv.Atoi(middlewareParams)
	if err != nil {
		app.errJSON(w, errors.New("invalid user"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errJSON(w, err)
		return
	}

	var prediction models.Prediction

	// verify the user edits his own post
	// idToDel, _ := strconv.Atoi(payload.ID)
	// if idUser != app.models.DB.DbUserIdOf(idToDel) {
	// 	app.errJSON(w, errors.New("not your post"))
	// } else {

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.GetSec(id)
		prediction = *m
		prediction.UpdatedAt = time.Now()
	}

	prediction.ID, _ = strconv.Atoi(payload.ID)
	prediction.Title = payload.Title
	prediction.Description = payload.Description
	prediction.CreatedAt = time.Now()
	prediction.UpdatedAt = time.Now()
	// prediction.Author, _ = strconv.Atoi(payload.Author)
	prediction.Author = idUser

	if prediction.ID == 0 {
		err = app.models.DB.InsertPrediction(prediction)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"predis_un\"") {
				app.errJSON(w, errors.New("title already exists, choose another one"))
			} else {
				app.errJSON(w, err)
			}

			return
		}
	} else {
		err = app.models.DB.EditPrediction(prediction)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"predis_un\"") {
				app.errJSON(w, errors.New("title already exists, choose another one"))
			} else {
				app.errJSON(w, err)
			}
			return
		}
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.errJSON(w, err)
		return
	}
	// }
}

func (app *application) deletePrediction(w http.ResponseWriter, r *http.Request) {

	// user id from request
	middlewareParams := r.Context().Value(UserParam{}).(string)
	idUser, err := strconv.Atoi(middlewareParams)

	if err != nil {
		app.logger.Print(errors.New("invalid uid parameter"))
		app.errJSON(w, err)
		return
	}

	// find the param in url
	delId := strings.Split(r.URL.String(), "/")
	idDeleteString := delId[len(delId)-1]

	idDelete, err := strconv.Atoi(idDeleteString)
	if err != nil {

		app.errJSON(w, err)
		return
	}

	if idUser == app.models.DB.DbUserIdOf(idDelete) {
		err = app.models.DB.DeletePrediction(idDelete)
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
	} else {
		app.errJSON(w, errors.New("not your post "))
	}

}

// handler for get MyAll
func (app *application) getAllMyPredictions(w http.ResponseWriter, r *http.Request) {

	// usr := r.Context().Value(UserParam{})
	// params := fmt.Sprintf("%v", usr)
	// id := fmt.Sprintf("%d", )
	// app.logger.Println(id)
	// app.logger.Printf("%T", id)

	//////////////////////////////////////////////////////////

	// get UserId param from Middleware
	// middlewareParams := r.Context().Value("idUser")
	//

	// idUser, err := strconv.Atoi(id)

	// if err != nil {
	// 	app.logger.Print(errors.New("invalid uid parameter"))
	// 	app.errJSON(w, err)
	// 	return
	// }

	/////////////////////////////////////////////////////////

	middlewareParams := r.Context().Value(UserParam{}).(string)
	idUser, err := strconv.Atoi(middlewareParams)

	if err != nil {
		app.logger.Print(errors.New("invalid uid parameter"))
		app.errJSON(w, err)
		return
	}

	predictions, err := app.models.DB.MyAll(idUser)

	if err != nil {
		app.errJSON(w, err)
		return
	}

	if len(predictions) == 0 {
		err = app.writeJSON(w, 204, "no items available", "predictions")
		if err != nil {
			app.errJSON(w, err)
			return
		}
	} else {
		err = app.writeJSON(w, http.StatusOK, predictions, "predictions")
		if err != nil {
			app.errJSON(w, err)
			return
		}
	}

}

type SomePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GET ONE PREDICTION
func (app *application) getOnePrediction(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errJSON(w, err)
		return
	}

	prediction_1, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, prediction_1, "prediction")
	if err != nil {
		app.errJSON(w, err)
		return
	}
}
