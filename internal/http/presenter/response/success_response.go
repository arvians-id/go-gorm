package response

import (
	"encoding/json"
	"github.com/arvians-id/go-gorm/internal/http/presenter/request"
	"net/http"
)

func ReturnSuccessOK(w http.ResponseWriter, status string, data interface{}) {
	var response WebResponse
	response.Code = http.StatusOK
	response.Status = status
	response.Data = data

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}

func ReturnSuccessPagesOK(w http.ResponseWriter, status string, data interface{}, pages request.PaginationData) {
	var response WebResponsePages
	response.Code = http.StatusOK
	response.Status = status
	response.Data = data
	response.Pages = pages

	marshal, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
