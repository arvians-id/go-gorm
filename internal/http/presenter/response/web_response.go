package response

import "github.com/arvians-id/go-gorm/internal/http/presenter/request"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponsePages struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Data   interface{}            `json:"data"`
	Pages  request.PaginationData `json:"pages"`
}
