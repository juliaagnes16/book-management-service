package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok || uname != "admin" || pwd != "admin" {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak boleh kosong atau salah"})
			c.Abort()
			return
		}
		c.Next()
	}
}
