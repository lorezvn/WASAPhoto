package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// If the user does not exist, it will be created, and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user User

	w.Header().Set("content-type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rt.baseLogger.WithError(err).Warning("Decoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := user.Username

	userID, _ := rt.db.CreateUser(username)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(userID); err != nil {
		rt.baseLogger.WithError(err).Warning("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
