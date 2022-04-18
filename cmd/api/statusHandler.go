package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	// convert struct to json using MARSHAL indent
	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) infoHandler(w http.ResponseWriter, r *http.Request) {
	currentInfo := InfoStatus{
		Developer:    "Mircea Brisan",
		Technologies: "ReactJS + Golang",
	}

	// convert struct to json using MARSHAL indent
	js, err := json.MarshalIndent(currentInfo, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
