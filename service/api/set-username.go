package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
Given the user id of the user who wants to set a new username, it sets the
the new username if the user exists
*/
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user User

	userToken := extractBearerToken(r)

	if !validToken(userToken) {
		rt.baseLogger.Error("Invalid Token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting userID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rt.baseLogger.WithError(err).Error("Decoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUsername := user.Username

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(userID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userToken != userID {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !validUsername(newUsername) {
		rt.baseLogger.Error("Invalid username")
		writeJSONResponse(w, http.StatusBadRequest, "Invalid username: a valid username should have 3-16 characters")
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rt.db.UsernameExists(newUsername) {
		rt.baseLogger.Error("Username chosen already exists")
		writeJSONResponse(w, http.StatusConflict, "Username chosen already exists")
		// w.WriteHeader(http.StatusConflict)
		return
	}

	if err := rt.db.ChangeUsername(userToken, newUsername); err != nil {
		rt.baseLogger.WithError(err).Error("Error updating username into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.baseLogger.Info("Username successfully set")
	w.WriteHeader(http.StatusNoContent)
}
