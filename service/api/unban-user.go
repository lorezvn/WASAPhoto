package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/database"
)

/*
	Given the id of the user who wants to unban another one (userID) and the
	id of the user to unban (banID), this method allows the unban operation
	(only if the user was previously in the ban status), meaning that the
	ex-banned user will be able to re-visualize the user profile of
	the user who previously banned him
*/

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
		rt.baseLogger.Error("Users can't unban themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = rt.db.UnbanUser(userToken, banID); err != nil {
		if errors.Is(err, database.ErrBanNotFound) {
			rt.baseLogger.WithError(err).Error("Error removing ban from DB")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			rt.baseLogger.WithError(err).Error("Error removing ban from DB")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rt.baseLogger.Info("User successfully unbanned, he will be able to re-visualize your profile")
	w.WriteHeader(http.StatusNoContent)
}
