package http

import (
	"mercadolibre/internal/location"
	"mercadolibre/internal/middleware"

	"github.com/labstack/echo/v4"
)

// Map location routes
func MapLocationRoutes(locationGroup *echo.Group, h location.Handlers, mw *middleware.MiddlewareManager) {
	locationGroup.POST("/topsecret", h.GetLocationBySatellites())
	locationGroup.POST("/topsecret_split/:satellite_name", h.PostTopSecretSplit())
	locationGroup.GET("/topsecret_split", h.GetTopSecretSplit())
}
