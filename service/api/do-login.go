package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*
If the user does not exist, it will be created, and an identifier is returned.
If the user exists, the user identifier is returned.
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rt.baseLogger.WithError(err).Error("Decoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := user.Username

	if !validUsername(username) {
		rt.baseLogger.Error("Invalid username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, _ := rt.db.CreateUser(username)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(userID); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
