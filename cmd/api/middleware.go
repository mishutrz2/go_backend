package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

// cross origins requests
func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// set content type header to be something that is permited
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(w, r)
	})
}

// check token
func (app *application) checkToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Vary", "Authorization")
		authHeader := r.Header.Get("Authorization")

		// if authHeader == "" {
		// 	// anonymous user maybe??
		// }

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errJSON(w, errors.New("invalid auth header"))
			return
		}

		if headerParts[0] != "Bearer" {
			app.errJSON(w, errors.New("unauthorized - no bearer"))
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))

		if err != nil {
			app.errJSON(w, errors.New("unauthorized - failed hmac check"))
		}

		if !claims.Valid(time.Now()) {
			app.errJSON(w, errors.New("unauthorized - token expired"))
		}

		if !claims.AcceptAudience("ceva.com") {
			app.errJSON(w, errors.New("unauthorized - invalid audience"))
		}

		if claims.Issuer != "ceva.com" {
			app.errJSON(w, errors.New("unauthorized - invalid issuer"))
		}

		userId, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.errJSON(w, errors.New("user id problem"))
		}

		log.Println("Middleware // Valid user: ", userId)

		// send context with userId via UserParam interface

		x := context.WithValue(context.Background(), UserParam{}, fmt.Sprint(userId))
		req := r.WithContext(x)

		next.ServeHTTP(w, req)
	})
}

type UserParam struct{}
