package httpserver

type HttpRequest struct {
	Body   []byte
	Params map[string]string
}

type HttpResponse struct {
	StatusCode int
	Body       any
}

func JSON(statusCode int, body any) HttpResponse {
	return HttpResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

type ControllerFunc func(HttpRequest) HttpResponse

type HttpServer interface {
	RegisterRoute(string, string, ControllerFunc)
	Start(string) error
}
