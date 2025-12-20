package web

import (
	"encoding/json"
	"net/http"
	"order/application/command"
	"order/application/query"
	httpserver "order/infra/http_server"

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

func (c *OrderController) PlaceOrder(req httpserver.HttpRequest) httpserver.HttpResponse {
	var input command.PlaceOrderInput
	if err := json.Unmarshal(req.Body, &input); err != nil {
		return httpserver.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	input.ID = uuid.NewString()
	if err := c.placeOrder.Run(input); err != nil {
		return httpserver.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return httpserver.JSON(http.StatusAccepted, map[string]string{
		"message": "your order is being processed",
		"id":      input.ID,
	},
	)
}

func (c *OrderController) GetOrderByID(req httpserver.HttpRequest) httpserver.HttpResponse {
	id := req.Params["id"]
	order, err := c.getOrderByID.Run(id)
	if err != nil {
		return httpserver.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return httpserver.JSON(http.StatusFound, order)
}
