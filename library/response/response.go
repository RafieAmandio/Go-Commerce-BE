package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponse(rw http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	// set header for json type
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(resp)
}
