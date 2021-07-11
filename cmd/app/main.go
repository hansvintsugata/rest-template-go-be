package main

import (
	"fmt"
	"net/http"

	"github.com/rest-template-go-be/internal/wire"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting service...")

	rootRouter := chi.NewRouter()
	rootRouter.Use(middleware.Recoverer)

	registerHandler(rootRouter)

	addr := fmt.Sprintf(":%d", 9000) //TODO
	logrus.WithField("addr", addr).Infof("Starting HTTP server")

	if err := http.ListenAndServe(addr, rootRouter); err != nil {
		logrus.Error(err)
	}
}

// registerHandler register all http handler
func registerHandler(router *chi.Mux) {
	router.Route("/", func(r chi.Router) {
		wire.InitializeHTTP().CreateHandler(router)
	})

	router.Get("/application/health", func(writer http.ResponseWriter, request *http.Request) {
		// TODO: healthcheck
	})

	sh := http.StripPrefix("/docs", http.FileServer(http.Dir("docs/swaggerui")))
	router.Get("/docs/*", func(writer http.ResponseWriter, request *http.Request) {
		sh.ServeHTTP(writer, request)
	})
}
