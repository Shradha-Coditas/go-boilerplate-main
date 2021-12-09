package server

import (
	"net/http"

	deleteAlbum "boilerplate/controller/albums/delete"
	getAlbum "boilerplate/controller/albums/get"
	patchAlbum "boilerplate/controller/albums/patch"
	postAlbum "boilerplate/controller/albums/post"
	postScrip "boilerplate/controller/scrip/post"
	postTagredis "boilerplate/controller/tagredis/post"
	postTags "boilerplate/controller/tags/post"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Group user related routes together
	v1Routes := router.Group("v1")
	{
		// Handle the GET requests at /v1/status
		v1Routes.GET("/status", func(c *gin.Context) {
			// send success response
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
			})
		})

		// Handle the Post requests at /v1/album
		v1Routes.POST("/album", postAlbum.Handler)

		// Handle the PATCH requests at /v1/album
		v1Routes.PATCH("/album", patchAlbum.Handler)

		// Handle the GET requests at /v1/album
		v1Routes.GET("/album", getAlbum.Handler)

		// Handle the DELETE requests at /v1/album
		v1Routes.DELETE("/album", deleteAlbum.Handler)

		// Handle the Post requests at /v1/tags
		v1Routes.POST("/tags", postTags.Handler)

		// Handle the Post requests at /v1/scrip
		v1Routes.POST("/scrip", postScrip.Handler)

		// Handle the Post requests at /v1/tagredis
		v1Routes.POST("/tagredis", postTagredis.Handler)

	}

	// Group user related routes together
	v2Routes := router.Group("v2")
	{
		// Handle the GET requests at /v2/status
		v2Routes.GET("/status", func(c *gin.Context) {
			// send success response
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
			})
		})

		// Handle the Post requests at /v2/album
		v2Routes.POST("/album", postAlbum.Handler)

	}
}
