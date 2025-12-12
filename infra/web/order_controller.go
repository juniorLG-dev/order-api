package web

import (
	"net/http"
	"order/application/command"
	"order/application/query"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	placeOrder   command.PlaceOrder
	getOrderByID query.GetOrderByID
}

func NewOrderController(
	placeOrder command.PlaceOrder,
	getOrderByID query.GetOrderByID,
) *OrderController {
	return &OrderController{
		placeOrder:   placeOrder,
		getOrderByID: getOrderByID,
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
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "your order is being processed",
		"id":      input.ID,
	},
	)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := c.getOrderByID.Run(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusFound, order)
}
