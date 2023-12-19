package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	var user User

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Error getting userID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rt.baseLogger.WithError(err).Warning("Decoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUsername := user.Username

	if err := rt.db.ChangeUsername(userID, newUsername); err != nil {
		http.Error(w, "Error updating username", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
