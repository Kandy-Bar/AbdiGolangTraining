package main

import (
	"20240606/middleware"
	"20240606/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//instalasi router gin

	gin.SetMode(gin.ReleaseMode) // set jadi realseseMode Jika Production atau UAT

	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	router.SetupRouter(r)

	log.Println("Running server on port 8080")

	r.Run(":8080")
}
