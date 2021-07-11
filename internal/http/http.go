package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rest-template-go-be/internal/http/hello/port/genhttp"
)

type HTTP struct {
}

var _ genhttp.ServerInterface = (*HTTP)(nil)

func ProvideHTTP() HTTP {
	return HTTP{}
}

func (h HTTP) CreateHandler(router chi.Router) http.Handler {
	var middlewares []genhttp.MiddlewareFunc
	return genhttp.HandlerFromMuxWithMiddleware(h, router, middlewares)
}

func (H HTTP) GetHelloWorld(w http.ResponseWriter, r *http.Request, params genhttp.GetHelloWorldParams) {
	panic("implement me")
}
