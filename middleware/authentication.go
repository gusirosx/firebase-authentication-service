package middleware

import (
	"firebase/service"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// Create an unexported global variable to hold the firebase connection pool.
var client *auth.Client = service.FirebaseInstance()

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
			return
		}

		_, err := client.VerifyIDToken(ctx, clientToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
}
