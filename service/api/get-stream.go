package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// If the user does not exist, it will be created, and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Error getting userID", http.StatusBadRequest)
		return
	}

	stream, err := rt.db.GetStream(userID)
	if err != nil {
		http.Error(w, "Error obtaining stream of photos", http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(stream); err != nil {
		rt.baseLogger.WithError(err).Warning("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
