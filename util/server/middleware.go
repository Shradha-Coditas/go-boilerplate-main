package server

import (
	"boilerplate/util/logwrapper"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// LanHeader - is for headers without auth.
type LanHeader struct {
	Language string `header:"language" binding:"required"`
	Platform int    `header:"platform" binding:"required"`
}

// AuthLanHeader - is for headers with auth.
type AuthLanHeader struct {
	Authorization string `header:"authorization" binding:"required"`
	Language      string `header:"language" binding:"required"`
	Platform      int    `header:"platform" binding:"required"`
}

// This middleware sets whether the user is logged in or not.
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Set("validUser", false)
		} else {
			logwrapper.Logger.Debugln("token : ", tokenString)
			var tokenData map[string]interface{}
			if json.Unmarshal([]byte(tokenString), &tokenData) == nil {
				c.Set("userData", tokenData)
				c.Set("validUser", true)
			} else {
				c.Set("validUser", false)
			}
		}
	}
}
