package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraiiian/mrt-schudules/modules/station"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	api := router.Group("/api/v1")

	station.Initiate(api)

	router.Run(":8080")
}
