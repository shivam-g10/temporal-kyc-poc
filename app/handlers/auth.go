package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

// User Sign up
func SignUp(w http.ResponseWriter, r *http.Request) {
	new_user := &NewUser{}
	err := json.NewDecoder(r.Body).Decode(new_user)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.WithError(err).Error("Something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = "Some Error Occurred"
		json.NewEncoder(w).Encode(resp)
	} else {
		user := &User{
			Email: new_user.Email,
			Id:    uuid.New().String(),
		}
		log.Infof("Created User %s", user.Id)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
