package main

import "github.com/gin-gonic/gin"

func main() {
	InitiateRouter()

	
}

fun InitiateRouter() {
	var router = gin.Default()
	 api = router.Group("/api/v1")

}

	station.Initiate(api)

	router.Run(":8080")
}