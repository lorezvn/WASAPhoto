package api

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

/*
	Given the id of the user who wants to unfollow another one (userID) and the
	id of the user who will be unfollowed (followID), if they are actually
	users of WASAPhoto and the user follows the other one, this method
	allows the unfollow action
*/

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
		rt.baseLogger.Error("Users can't unfollow themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rt.db.BanExists(followID, userToken) {
		rt.baseLogger.Error("This user was banned, impossible to perform the unfollow request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err = rt.db.UnfollowUser(userToken, followID); err != nil {
		if errors.Is(err, errors.New("Follow not found")) {
			rt.baseLogger.WithError(err).Error("Error inserting follow into DB")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			rt.baseLogger.WithError(err).Error("Error inserting follow into DB")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rt.baseLogger.Info("User successfully unfollowed")
	w.WriteHeader(http.StatusNoContent)
}
