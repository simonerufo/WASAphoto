package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	//login
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	rt.router.PUT("/profiles/:user_id/username", rt.wrap(rt.setMyUsername))
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

	rt.router.PUT("/profiles/:user_id/following/:following_id", rt.wrap(rt.followUser))
	rt.router.DELETE("/profiles/:user_id/following/:following_id", rt.wrap(rt.unfollowUser))
	rt.router.GET("/profiles/:user_id/follow", rt.wrap(rt.GetFollowers))

	rt.router.PUT("/profiles/:user_id/bans/:target_uid", rt.wrap(rt.BanUser))
	rt.router.DELETE("/profiles/:user_id/bans/:target_uid", rt.wrap(rt.UnbanUser))
	rt.router.GET("/profiles/:user_id/ban", rt.wrap(rt.GetBan))

	rt.router.GET("/profiles/:user_id/stream", rt.GetMyStream)
	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
