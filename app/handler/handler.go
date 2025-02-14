package handler

import (
	"encoding/json"
	"net/http"
	"splitto/app/model"
)

type Handler struct {
	DB *model.Model
}

func WriteResponse(res map[string]string, code int, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
