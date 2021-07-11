//+build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/rest-template-go-be/internal/http"
	"github.com/rest-template-go-be/internal/service"
)

func InitializeHTTP() http.HTTP {
	panic(wire.Build(
		http.ProvideHTTP,
		service.ProvideHelloService,
	))
}
