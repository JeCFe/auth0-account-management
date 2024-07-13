package main

import (
	"jecfe/auth0-account-management/controller"
	"jecfe/auth0-account-management/middleware"
	_ "jecfe/auth0-account-management/middleware"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"

	"github.com/gin-gonic/gin"
)

func GetClaimsFromContext(ctx *gin.Context) validator.RegisteredClaims {
	claims, ok := ctx.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}

	return claims.RegisteredClaims
}

func main() {
	router := gin.Default()
	x := controller.NewController();
	router.Use(middleware.CheckJWT())

	x.
	
	router.GET("/", func(ctx *gin.Context) {
		claims := GetClaimsFromContext(ctx)
		ctx.JSON(http.StatusOK, claims.Subject)
	})

	router.Run(":5080")
}
