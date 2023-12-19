package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Error getting userID", http.StatusBadRequest)
		return
	}

	followID, err := strconv.Atoi(ps.ByName("followID"))
	if err != nil {
		http.Error(w, "Error getting followID", http.StatusBadRequest)
		return
	}

	if err = rt.db.FollowUser(userID, followID); err != nil {
		http.Error(w, "Error following user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
