package web

import "github.com/gin-gonic/gin"

type ControllerGroup interface {
	PlaceOrder(ctx *gin.Context)
}
