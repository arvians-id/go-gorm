package response

import (
	"encoding/json"
	"net/http"
)

var (
	ErrorNotFound         = "sql: no rows in result set"
	WrongPassword         = "wrong password"
	VerificationExpired   = "reset password verification is expired"
	ResponseErrorNotFound = "data not found"
)

func ReturnErrorNotFound(w http.ResponseWriter, err error, data interface{}) {
	var response WebResponse
	response.Code = http.StatusNotFound
	response.Status = "data not found"
	response.Data = data

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(marshal)
}

func ReturnErrorInternalServerError(w http.ResponseWriter, err error, data interface{}) {
	var response WebResponse
	response.Code = http.StatusInternalServerError
	response.Status = err.Error()
	response.Data = data

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(marshal)
}

func ReturnErrorBadRequest(w http.ResponseWriter, err error, data interface{}) {
	var response WebResponse
	response.Code = http.StatusBadRequest
	response.Status = err.Error()
	response.Data = data

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(marshal)
}

func ReturnErrorUnauthorized(w http.ResponseWriter, err error, data interface{}) {
	var response WebResponse
	response.Code = http.StatusUnauthorized
	response.Status = err.Error()
	response.Data = data

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(marshal)
}
