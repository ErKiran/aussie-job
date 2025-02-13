package controllers

import (
	"log"
	"time"

	"aussie-jobs/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	db, err := repositories.GetDB()
	if err != nil {
		log.Fatal("unable to get db", err)
	}
	controller := NewController(db)
	r := gin.Default()
	r.UseRawPath = true
	r.UnescapePathValues = false

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS, HEAD, PATCH"},
		AllowHeaders:     []string{"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// setting the swagger host static to the development api server
	// because we don't want to expose the api documents to public
	// as the swagger is not protected by any authentication and so the exposed api's
	// docs.SwaggerInfo.Host = "http:localhost:8080/"
	// docs.SwaggerInfo.Schemes = []string{"https", "http"}
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	job := r.Group("api/v1")
	{
		job.GET("/search", controller.SeekController.SearchJob)
		job.GET("/analysis/location", controller.SeekController.AnalyzeJobs)
	}

	return r
}
