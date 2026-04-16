package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	stationGroup := router.Group("/stations")
	stationGroup.GET("", func(c *gin.Context) {
		GetAllStations(c, stationService)
	})

	stationGroup.GET("/:id", func(c *gin.Context) {
		CheckScheduleByStationID(c, stationService)
	})
}

func GetAllStations(c *gin.Context, service *Service) {
	data, err := service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "successfully get all stations",
		"data":    data,
	})
}

func CheckScheduleByStationID(c *gin.Context, service *Service) {
	id := c.Param("id")

	datas, err := service.GetScheduleByStationID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "successfully get schedules by station",
		"data":    datas,
	})
}
