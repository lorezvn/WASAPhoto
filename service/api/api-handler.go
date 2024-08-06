package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)
	rt.router.GET("/users/", rt.searchUsers)
	rt.router.PUT("/users/:userID/username", rt.setMyUsername)
	rt.router.PUT("/users/:userID/follow/:followID", rt.followUser)
	rt.router.DELETE("/users/:userID/follow/:followID", rt.unfollowUser)
	rt.router.POST("/users/:userID/photos/", rt.uploadPhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID", rt.deletePhoto)
	rt.router.POST("/users/:userID/photos/:photoID/comments/", rt.commentPhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID/comments/:commentID", rt.uncommentPhoto)
	rt.router.PUT("/users/:userID/photos/:photoID/likes/:likeID", rt.likePhoto)
	rt.router.DELETE("/users/:userID/photos/:photoID/likes/:likeID", rt.unlikePhoto)
	rt.router.GET("/users/:userID/ban/", rt.getBannedUsers);
	rt.router.PUT("/users/:userID/ban/:banID", rt.banUser)
	rt.router.DELETE("/users/:userID/ban/:banID", rt.unbanUser)
	rt.router.GET("/users/:userID/stream", rt.getMyStream)
	rt.router.GET("/users/:userID/profile", rt.getUserProfile)

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
