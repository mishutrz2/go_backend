package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// helper function
func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// var x SomePayload
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// here we manage our routes
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// chain for routes
	secure := alice.New(app.checkToken)

	// Status and Info handlers
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/info", app.infoHandler)

	// register a new account
	router.HandlerFunc(http.MethodPost, "/v1/signup", app.registerAccount)

	// handler jwt
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	// get Prediction/s handlers
	router.HandlerFunc(http.MethodGet, "/v1/all/:id", app.getOnePrediction)
	router.HandlerFunc(http.MethodGet, "/v1/all", app.getAllPredictions)

	router.GET("/v1/my", app.wrap(secure.ThenFunc(app.getAllMyPredictions)))
	router.HandlerFunc(http.MethodGet, "/v1/my/:id", app.getOnePrediction)

	router.POST("/v1/admin/edit", app.wrap(secure.ThenFunc(app.editPrediction)))
	// router.HandlerFunc(http.MethodPost, "/v1/admin/edit", app.editPrediction)

	router.GET("/v1/admin/delete/:id", app.wrap(secure.ThenFunc(app.deletePrediction)))
	// router.HandlerFunc(http.MethodGet, "/v1/admin/delete/:id", app.deletePrediction)

	return app.enableCORS(router)
}
