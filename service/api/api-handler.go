package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	//login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	//user actions
	rt.router.PUT("/profiles/:user_id/username", rt.wrap(rt.setMyUsername))
	rt.router.PUT("/profiles/:user_id/followed/:target_id", rt.wrap(rt.followUser))
	//getting user profile
	rt.router.GET("/profiles/:user_id/username",rt.wrap(rt.getProfileByUsername))
	//search
	rt.router.GET("/profiles/:user_id/profile", rt.wrap(rt.getProfile))
	rt.router.POST("/profiles/:user_id/profile", rt.wrap(rt.uploadPhoto))

	rt.router.GET("/profiles/:user_id/photos/:photo_id", rt.wrap(rt.getPhoto))
	rt.router.DELETE("/profiles/:user_id/photos/:photo_id",rt.wrap(rt.deletePhoto))

	rt.router.PUT("/profiles/:user_id/likes/:photo_id", rt.wrap(rt.LikePhoto))
	rt.router.GET("/photos/:photo_id/likes", rt.wrap(rt.GetPhotoLikes))
	rt.router.DELETE("/profiles/:user_id/likes/:photo_id", rt.wrap(rt.RemoveLike))

	rt.router.POST("/profiles/:user_id/comments/:photo_id", rt.wrap(rt.CommentPhoto))
	rt.router.GET("/profiles/:user_id/photos/:photo_id/comments", rt.wrap(rt.GetPhotoComments))
	rt.router.DELETE("/profiles/:user_id/comments/:comment_id", rt.wrap(rt.RemoveComment))

	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
