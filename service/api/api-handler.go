package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)
	rt.router.PUT("/users/:userID/username", rt.setMyUsername)
	rt.router.PUT("/users/:userID/follow/:followID", rt.followUser)
	rt.router.POST("/users/:userID/photos/", rt.uploadPhoto)
	rt.router.GET("/users/:userID/photos/", rt.getMyStream)
	//rt.router.DELETE("/users/:userID/photos/:photoID", rt.deletePhoto)

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}