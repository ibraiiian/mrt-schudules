package main

import "github.com/gin-gonic/gin"

func main() {
	InitiateRouter()

	
}

fun InitiateRouter() {
	var router = gin.Default()
	var api = router.Group("/api/v1")

	router.Run(":8080")
}