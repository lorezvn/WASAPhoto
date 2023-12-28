package api

import (
	"net/http"
	"strconv"
	"errors"
	"github.com/julienschmidt/httprouter"
)

/*
	Given the photo id and the id of the User who liked the Photo (likeID), deletes the Like with the user id 
	that corresponds to the one given
*/

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userToken := extractBearerToken(r)

	if !validToken(userToken) {
		rt.baseLogger.Error("Invalid Token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoAuthorID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting userID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoID, err := strconv.Atoi(ps.ByName("photoID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting photoID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(ps.ByName("likeID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting likeID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(photoAuthorID) || !rt.db.UserExists(userID){
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !rt.db.PhotoExists(photoID, photoAuthorID) {
		rt.baseLogger.Error("Photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if (userToken != userID) {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := rt.db.DeleteLike(userID, photoID); err != nil {
		if errors.Is(err, errors.New("Like not found")) {
			rt.baseLogger.WithError(err).Error("Error removing like from DB")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			rt.baseLogger.WithError(err).Error("Error removing like from DB")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rt.baseLogger.Info("Removed successfully a like from the photo")
	w.WriteHeader(http.StatusNoContent)
}
