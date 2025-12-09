package web

import (
	"net/http"
	"order/application/command"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	placeOrder command.PlaceOrder
}

func NewOrderController(
	placeOrder command.PlaceOrder,
) *OrderController {
	return &OrderController{
		placeOrder: placeOrder,
	}
}

func (c *OrderController) PlaceOrder(ctx *gin.Context) {
	var input command.PlaceOrderInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	input.ID = uuid.NewString()
	if err := c.placeOrder.Run(input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "your order is being processed"})
}
