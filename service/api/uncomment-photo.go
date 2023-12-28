package api

import (
	"net/http"
	"strconv"
	"errors"
	"github.com/julienschmidt/httprouter"
)

/*
	Given the photo id and a comment id, deletes the Comment with the comment id 
	that corresponds to the one given
*/
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	commentID, err := strconv.Atoi(ps.ByName("commentID"))
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting commentID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if !rt.db.UserExists(userToken) || !rt.db.UserExists(photoAuthorID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !rt.db.PhotoExists(photoID, photoAuthorID) {
		rt.baseLogger.Error("Photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := rt.db.DeleteComment(commentID); err != nil {
		if errors.Is(err, errors.New("Comment not found")) {
			rt.baseLogger.WithError(err).Error("Error removing comment from DB")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			rt.baseLogger.WithError(err).Error("Error removing comment from DB")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rt.baseLogger.Info("Comment successfully removed")
	w.WriteHeader(http.StatusNoContent)
}

