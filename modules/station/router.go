package station station

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {
	stationService := station.NewService()

	stationGroup := router.Group("/stations")
	stationGroup.GET "",func(c *gin.Context) {
		GetAllStations(c, stationService)
	})
}
 func  GetAllStations(c *gin.Context, service Service) {) {
	data, err := service.GetAllStations()
	if err != nil {
		c.JSON (code: https, StatusBadRequest,
			response: response.ApiResponse{
				success: false,
				messege: err.Error(),
				data: nil,
			})
		

	}

	c.JSON(
		http.StatusOK,
		response.ApiResponse{}
		success: true,
		messege: "successfully get all stations",
		data: data,
	)
}

	