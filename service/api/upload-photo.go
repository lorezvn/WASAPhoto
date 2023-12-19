package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

// If the user does not exist, it will be created, and an identifier is returned.
// If the user exists, the user identifier is returned.
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Error getting userID", http.StatusBadRequest)
		return
	}

	image, err := io.ReadAll(r.Body)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Reading image data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoID, date, err := rt.db.InsertPhoto(userID, image)
	if err != nil {
		http.Error(w, "Error inserting photo", http.StatusBadRequest)
		return
	}

	photo := Photo{
		PhotoID:  photoID,
		UserID:   userID,
		Image:    image,
		Date:     date,
		Likes:    []Like{},
		Comments: []Comment{},
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(photo); err != nil {
		rt.baseLogger.WithError(err).Warning("Encoding JSON failed")
		w.WriteHeader(http.StatusBadRequest)
	}
}
