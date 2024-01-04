package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Given the user id, returns the stream of photos
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

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

	stream, err := rt.db.GetStream(userID)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting stream of photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(stream); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rt.baseLogger.Info("Stream of photos returned")
}
