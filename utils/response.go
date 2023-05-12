package utils

import (
	"encoding/json"
	"entdemo-api/model"
	"net/http"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {
	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
