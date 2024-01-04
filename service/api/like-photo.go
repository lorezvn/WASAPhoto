package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

/*
	Given the photo id, and the id of the User who wants to add a
	like to the Photo (likeID), it adds a like
*/
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

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

	if !rt.db.UserExists(userToken) || !rt.db.UserExists(photoAuthorID) || !rt.db.UserExists(userID) {
		rt.baseLogger.Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !rt.db.PhotoExists(photoID, photoAuthorID) {
		rt.baseLogger.Error("Photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userToken != userID {
		rt.baseLogger.Error("Access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if userToken == photoAuthorID {
		rt.baseLogger.Error("Users can't like their own photos")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rt.db.BanExists(photoAuthorID, userToken) {
		rt.baseLogger.Error("This user was banned, impossible to perform the like request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	date, err := rt.db.InsertLike(userToken, photoID)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error inserting like into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	like := Like{
		UserID:  userToken,
		PhotoID: photoID,
		Date:    date,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(like); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
	}

	rt.baseLogger.Info("Like successfully added")
}
