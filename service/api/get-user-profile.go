package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Get user profile by user ID
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	userToken := extractBearerToken(r)

	if !validToken(userToken) {
		rt.baseLogger.Error("Invalid Token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error getting userID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(userID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userProfile, err := rt.db.GetUserProfile(userID)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting user profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(userProfile); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rt.baseLogger.Info("User profile returned")
}
