package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Returns all the users matching the username query parameter
func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	userToken := extractBearerToken(r)

	if !validToken(userToken) {
		rt.baseLogger.Error("Invalid Token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !rt.db.UserExists(userToken) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usernameQuery := r.URL.Query().Get("username")

	// Check if usernameQuery is empty
	if usernameQuery == "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]User{})
		return
	}

	users, err := rt.db.SearchUsers(usernameQuery)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error searching users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(users); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
