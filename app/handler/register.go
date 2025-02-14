package handler

import (
	"encoding/json"
	"net/http"
	"splitto/app/model"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Extract User Fields
	newUser, err := getUserFields(w, r)
	if err != nil {
		res := map[string]string{
			"message": "Invalid Parameters",
		}
		WriteResponse(res, http.StatusBadRequest, w, r)
		return
	}

	// Save User to DB
	userID, err := h.DB.AddNewUser(newUser)

	// User Created, Return user_id
	res := map[string]string{
		"message": "User Registered Successfully",
		"user_id": userID,
	}
	WriteResponse(res, http.StatusCreated, w, r)
	return
}

func getUserFields(w http.ResponseWriter, r *http.Request) (model.UserRegister, error) {
	var newUser model.UserRegister

	err := json.NewDecoder(r.Body).Decode(&newUser)
	return newUser, err
}
