package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	//login
	rt.router.POST("/session", rt.doLogin)

	//username set/change
	rt.router.PUT("/profiles/:user_id/username", rt.setMyUsername)
	//getting user profile
	rt.router.GET("/profile", rt.getProfileByUsername)
	rt.router.GET("/profiles/:user_id/profile", rt.getProfile)
	//photo utility
	rt.router.POST("/profiles/:user_id/profile", rt.uploadPhoto)
	rt.router.GET("/profiles/:user_id/photos/:photo_id", rt.getPhoto)
	rt.router.DELETE("/profiles/:user_id/photos/:photo_id",rt.deletePhoto)
	//photo likes
	rt.router.PUT("/profiles/:user_id/likes/:photo_id", rt.LikePhoto)
	rt.router.GET("/photos/:photo_id/likes", rt.GetPhotoLikes)
	rt.router.DELETE("/profiles/:user_id/likes/:photo_id", rt.RemoveLike)
	//photo comments
	rt.router.POST("/profiles/:user_id/comments/:photo_id", rt.CommentPhoto)
	rt.router.GET("/profiles/:user_id/photos/:photo_id/comments", rt.GetPhotoComments)
	rt.router.DELETE("/profiles/:user_id/comments/:comment_id", rt.RemoveComment)
	//follow
	rt.router.PUT("/profiles/:user_id/following/:following_id", rt.followUser)
	rt.router.DELETE("/profiles/:user_id/following/:following_id", rt.unfollowUser)
	rt.router.GET("/profiles/:user_id/followers", rt.GetFollowers)
	rt.router.GET("/profiles/:user_id/following", rt.GetFollowing)
	//ban
	rt.router.PUT("/profiles/:user_id/bans/:target_uid", rt.BanUser)
	rt.router.DELETE("/profiles/:user_id/bans/:target_uid", rt.UnbanUser)
	rt.router.GET("/profiles/:user_id/bans/:target_uid", rt.GetBan)
	rt.router.GET("/profiles/:user_id/ban", rt.GetBanList)

	//profile stream
	rt.router.GET("/profiles/:user_id/stream", rt.GetMyStream)
	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
