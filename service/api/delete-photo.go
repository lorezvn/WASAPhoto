package api

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

/*
Given the user id of the user that uploaded the photo and the photo id of the
photo itself, it deletes the photo from the photos uploaded by the user
*/
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(photoAuthorID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userToken != photoAuthorID {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := rt.db.DeletePhoto(photoID); err != nil {
		if errors.Is(err, errors.New("Photo not found")) {
			rt.baseLogger.WithError(err).Error("Error removing photo from DB")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			rt.baseLogger.WithError(err).Error("Error removing photo from DB")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rt.baseLogger.Info("Photo successfully removed")
	w.WriteHeader(http.StatusNoContent)
}
