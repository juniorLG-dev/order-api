package web

import (
	httpserver "order/infra/http_server"
)

type ControllerGroup interface {
	PlaceOrder(httpserver.HttpRequest) httpserver.HttpResponse
	GetOrderByID(httpserver.HttpRequest) httpserver.HttpResponse
}
