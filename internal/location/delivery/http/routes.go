package http

import (
	"mercadolibre/internal/location"
	"mercadolibre/internal/middleware"

	"github.com/labstack/echo/v4"
)

// Map news routes
func MapLocationRoutes(locationGroup *echo.Group, h location.Handlers, mw *middleware.MiddlewareManager) {
	locationGroup.GET("/topsecret", h.GetLocationBySatellites())
}
