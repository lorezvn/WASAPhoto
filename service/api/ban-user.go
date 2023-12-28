package api

import (
	"net/http"
	"strconv"
	"errors"
	"github.com/julienschmidt/httprouter"
)

/*
	Given the id of the user who wants to ban another one (userID) and the
	id of the user to ban (banID), this method allows the ban operation, meaning
	that the banned user will no longer be able to visualize the user profile of
	the user who banned him
*/

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	banID, err := strconv.Atoi(ps.ByName("banID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting banID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(userID) || !rt.db.UserExists(banID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userToken != userID {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if userToken == banID {
		rt.baseLogger.Error("Users can't ban themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = rt.db.BanUser(userToken, banID); err != nil {
		rt.baseLogger.WithError(err).Error("Error inserting ban into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
		
	}

	rt.baseLogger.Info("User successfully banned, he will no longer be able to visualize your profile")
	w.WriteHeader(http.StatusNoContent)
}
