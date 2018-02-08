package utils

import (
	"encoding/json"
	"net/http"
)

// Response for standard response type
type Response struct {
	Status     string
	StatusCode int
	Message    string
	Data       interface{}
}

func SendAjaxResponseSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response := &Response{
		Status:     "ok",
		StatusCode: 200,
		Message:    "Success",
		Data:       data,
	}

	byt, _ := json.Marshal(response)
	w.Write(byt)
}

func SendAjaxResponseError(w http.ResponseWriter, Message string) {
	w.Header().Set("Content-Type", "application/json")

	response := &Response{
		Status:     "not ok",
		StatusCode: 502,
		Message:    Message,
		Data:       nil,
	}

	byt, _ := json.Marshal(response)
	w.Write(byt)
}
