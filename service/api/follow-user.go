package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
	Given the id of the user who wants to follow another one (userID) and the
	id of the user who will be followed (followID), if they are actually
	users of WASAPhoto, this method allows the follow action
*/

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	followID, err := strconv.Atoi(ps.ByName("followID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting followID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(userID) || !rt.db.UserExists(followID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userToken != userID {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if userToken == followID {
		rt.baseLogger.Error("Users can't follow themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rt.db.BanExists(followID, userToken) {
		rt.baseLogger.Error("This user was banned, impossible to perform the follow request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err = rt.db.FollowUser(userToken, followID); err != nil {
		rt.baseLogger.WithError(err).Error("Error inserting follow into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rt.baseLogger.Info("User successfully followed")
	w.WriteHeader(http.StatusNoContent)
}
