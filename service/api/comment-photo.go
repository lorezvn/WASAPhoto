package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Given the photo id and a message, it creates a new Comment
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	var comment Comment

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

	if !rt.db.PhotoExists(photoID, photoAuthorID) {
		rt.baseLogger.Error("Photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if rt.db.BanExists(photoAuthorID, userToken) {
		rt.baseLogger.Error("This user was banned, impossible to perform the comment request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		rt.baseLogger.WithError(err).Error("Decoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := comment.Message

	if !validMessage(message) {
		rt.baseLogger.Error("Invalid message")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, err := rt.db.GetUsernameByUserID(userToken)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentID, date, err := rt.db.InsertComment(userToken, photoID, message)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error inserting comment into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment = Comment{
		CommentID: commentID,
		UserID:    userToken,
		Username:  username,
		PhotoID:   photoID,
		Message:   message,
		Date:      date,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
	}

	rt.baseLogger.Info("Comment successfully posted")
}
