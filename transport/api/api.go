package api

import (
	_ "embed"
	"fmt"
	"net/http"
	transportApiv1 "newsletter_backend_api/transport/api/v1"

	"github.com/go-chi/chi"
	httpx "go.strv.io/net/http"

	"log/slog"
)

//go:embed v1/openapi.yaml
var OpenAPI []byte

type Controller struct {
	*chi.Mux

	service transportApiv1.Service
	version string
}

func NewController(
	service transportApiv1.Service,
	version string,
) (*Controller, error) {
	controller := &Controller{
		service: service,
		version: version,
	}
	controller.initRouter()
	return controller, nil
}

func (c *Controller) initRouter() {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// r.Use(httpx.LoggingMiddleware(util.NewServerLogger("httpx.LoggingMiddleware")))
		// r.Use(httpx.RecoverMiddleware(util.NewServerLogger("httpx.RecoverMiddleware").WithStackTrace(slog.Level)))
		// TODO: Add authentication middleware
		// authenticate := middleware.Authenticate(c.logger, c.tokenParser)

		v1Handler := transportApiv1.NewHandler(
			c.service,
		)

		r.Route("/api", func(r chi.Router) {
			r.Get("/openapi.yaml", c.OpenAPI)
			r.Mount("/v1", v1Handler)
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		slog.Info("ping triggered")
		w.WriteHeader(http.StatusNoContent)
	})

	r.Get("/version", c.Version)

	c.Mux = r
}

// TODO: Improve this handler.
func (c *Controller) OpenAPI(w http.ResponseWriter, _ *http.Request) {
	encodeFunc := func(w http.ResponseWriter, data any) error {
		d, ok := data.([]byte)
		if !ok {
			return fmt.Errorf("expected byte slice: got %T", data)
		}
		if _, err := w.Write(d); err != nil {
			return fmt.Errorf("writing openapi content: %w", err)
		}
		return nil
	}
	if err := httpx.WriteResponse(
		w,
		OpenAPI,
		http.StatusOK,
		httpx.WithEncodeFunc(encodeFunc),
		httpx.WithContentType(httpx.ApplicationYAML),
	); err != nil {
		slog.Error("writing response", slog.Any("error", err))
	}
}

func (c *Controller) Version(w http.ResponseWriter, _ *http.Request) {
	if err := httpx.WriteResponse(
		w,
		c.version,
		http.StatusOK,
	); err != nil {
		slog.Error("writing response", slog.Any("error", err))
	}
}
