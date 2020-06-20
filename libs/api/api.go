package api

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

type API struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(code, message string, data interface{}) *API {
	resp := API{}
	resp.Code = code
	resp.Message = message
	if data != nil || reflect.ValueOf(data).Kind() == reflect.Ptr {
		resp.Data = data
	}
	return &resp
}

func Write(w http.ResponseWriter, r *API) {
	js, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpCode, _ := strconv.Atoi(r.Code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(js)
}
