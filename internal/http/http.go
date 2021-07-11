package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rest-template-go-be/internal/http/hello/port/genhttp"
	"github.com/rest-template-go-be/internal/service"
	httputil "github.com/rest-template-go-be/pkg/http"
)

type HTTP struct {
	HelloService service.HelloService
}

var _ genhttp.ServerInterface = (*HTTP)(nil)

func ProvideHTTP(service service.HelloService) HTTP {
	return HTTP{
		HelloService: service,
	}
}

func (h HTTP) CreateHandler(router chi.Router) http.Handler {
	var middlewares []genhttp.MiddlewareFunc
	return genhttp.HandlerFromMuxWithMiddleware(h, router, middlewares)
}

func (h HTTP) GetHelloWorld(w http.ResponseWriter, r *http.Request, params genhttp.GetHelloWorldParams) {
	greeting, err := h.HelloService.Greeting(params.Flag)
	if err != nil {
		httputil.WriteResponse(w, "unexpected error", 500, nil, err)
		return
	}
	httputil.WriteResponse(w, "default.success", 200, greeting, nil)
}
