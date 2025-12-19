package adapter

import (
	"io"
	"net/http"
	httpserver "order/infra/http_server"

	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	gin *gin.Engine
}

func NewGinAdapter() *GinAdapter {
	return &GinAdapter{
		gin: gin.Default(),
	}
}

func (g *GinAdapter) AdaptController(handler httpserver.ControllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := make(map[string]string)
		for _, param := range ctx.Params {
			params[param.Key] = param.Value
		}
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		defer ctx.Request.Body.Close()
		request := httpserver.HttpRequest{
			Body:   bodyBytes,
			Params: params,
		}
		response := handler(request)

		if response.Body == nil {
			ctx.Status(response.StatusCode)
			return
		}
		ctx.JSON(response.StatusCode, response.Body)
	}
}

func (g *GinAdapter) RegisterRoute(method, path string, handler httpserver.ControllerFunc) {
	g.gin.Handle(method, path, g.AdaptController(handler))
}

func (g *GinAdapter) Start(port string) error {
	return g.gin.Run(port)
}
