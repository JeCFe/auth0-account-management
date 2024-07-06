package main

import (
	"jecfe/auth0-account-management/controller"
	"jecfe/auth0-account-management/docs"
	_ "jecfe/auth0-account-management/docs"

	_ "jecfe/auth0-account-management/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3020"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/v1")
	{
		example := v1.Group("/example")
		{
			example.GET("/ping", c.PingExample)	
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3020")
}