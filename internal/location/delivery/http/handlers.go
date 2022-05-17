package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/pkg/logger"
)

// Location handlers
type locationHandlers struct {
	cfg        *config.Config
	locationUC location.UseCase
	logger     logger.Logger
}

// NewNewsHandlers News handlers constructor
func NewLocationHandlers(cfg *config.Config, locationUC location.UseCase, logger logger.Logger) location.Handlers {
	return &locationHandlers{cfg: cfg, locationUC: locationUC, logger: logger}
}

// GetLocationBySatellites godoc
// @Summary Get by id news
// @Description Get by id news handler
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "news_id"
// @Success 200 {object} models.Location
// @Router /news/{id} [get]
func (h locationHandlers) GetLocationBySatellites() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, map[string]string{"status": "OK", "message": "Hola mundo"})
	}
}
