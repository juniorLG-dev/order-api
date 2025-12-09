package web

import "github.com/gin-gonic/gin"

func InitRoutes(rg *gin.RouterGroup, c ControllerGroup) {
	rg.POST("/order", c.PlaceOrder)
}
