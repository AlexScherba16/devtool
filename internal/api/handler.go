package api

import (
	"devtool/internal/services/devices"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

type handler struct {
	devHandler devices.DevicesHandler
}

func NewHandler() (handler, error) {
	devHandler, err := devices.NewDevicesHandler()
	if err != nil {
		return handler{}, err
	}

	return handler{
		devHandler: devHandler,
	}, nil
}

func (h *handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Timeout(10 * time.Second))

		r.Route("/devices", func(r chi.Router) {
			r.Get("/", h.devHandler.GetConnectedDevices)
		})
	})

	return router
}
