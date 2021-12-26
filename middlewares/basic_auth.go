package middlewares

import "github.com/gin-gonic/gin"

// Test Conceptual Middleware Auth Token
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"pragmatic": "reviews",
	})
}
