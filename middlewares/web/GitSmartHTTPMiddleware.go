package webMiddlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GitSmartHTTPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.UserAgent(), "git/") {
			c.String(http.StatusForbidden, "Access denied: Not a git client")
			c.Abort()
			return
		}

		if c.Query("service") != "git-upload-pack" && c.ContentType() != "application/x-git-upload-pack-request" {
			c.String(http.StatusForbidden, "Access denied: Not a smart HTTP request")
			c.Abort()
			return
		}
	}
}
