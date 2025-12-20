package web

import (
	httpserver "order/infra/http_server"
)

func InitRoutes(server httpserver.HttpServer, c ControllerGroup) {
	server.RegisterRoute("POST", "/order", c.PlaceOrder)
	server.RegisterRoute("GET", "/order/:id", c.GetOrderByID)
}
