package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// login
	rt.router.POST("/session", rt.doLogin)

	// username set/change
	rt.router.PUT("/profiles/:user_id/username", rt.setMyUserName)
	// getting user profile
	rt.router.GET("/profile", rt.getProfileByUsername)
	rt.router.GET("/profiles/:user_id/profile", rt.getUserProfile)
	// photo utility
	rt.router.POST("/profiles/:user_id/profile", rt.uploadPhoto)
	rt.router.GET("/profiles/:user_id/photos/:photo_id", rt.getPhoto)
	rt.router.DELETE("/profiles/:user_id/photos/:photo_id", rt.deletePhoto)
	// photo likes
	rt.router.PUT("/profiles/:user_id/likes/:photo_id", rt.likePhoto)
	rt.router.GET("/photos/:photo_id/likes", rt.getLikes)
	rt.router.DELETE("/profiles/:user_id/likes/:photo_id", rt.unlikePhoto)
	// photo comments
	rt.router.POST("/profiles/:user_id/comments/:photo_id", rt.commentPhoto)
	rt.router.GET("/profiles/:user_id/photos/:photo_id/comments", rt.getComments)
	rt.router.DELETE("/profiles/:user_id/comments/:photo_id", rt.uncommentPhoto)
	// follow
	rt.router.PUT("/profiles/:user_id/following/:following_id", rt.followUser)
	rt.router.DELETE("/profiles/:user_id/following/:following_id", rt.unfollowUser)
	rt.router.GET("/profiles/:user_id/followers", rt.getFollowers)
	rt.router.GET("/profiles/:user_id/following", rt.getFollowing)
	// ban
	rt.router.PUT("/profiles/:user_id/bans/:target_uid", rt.banUser)
	rt.router.DELETE("/profiles/:user_id/bans/:target_uid", rt.unbanUser)
	rt.router.GET("/profiles/:user_id/bans/:target_uid", rt.getBan)
	rt.router.GET("/profiles/:user_id/ban", rt.getBanList)

	// profile stream
	rt.router.GET("/profiles/:user_id/stream", rt.getMyStream)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
