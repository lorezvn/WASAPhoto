package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

// Takes an image and returns a new Photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	image, err := io.ReadAll(r.Body)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error reading image data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(image) == 0 {
		rt.baseLogger.Error("Invalid image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoID, date, err := rt.db.InsertPhoto(userToken, image)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error inserting photo into DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo := Photo{
		PhotoID:  photoID,
		UserID:   userToken,
		Image:    image,
		Date:     date,
		Likes:    []Like{},
		Comments: []Comment{},
	}

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(photo); err != nil {
		rt.baseLogger.WithError(err).Error("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rt.baseLogger.Info("Photo successfully uploaded")
}
