package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
	"mercadolibre/pkg/httpErrors"
	"mercadolibre/pkg/logger"
	"mercadolibre/pkg/utils"
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
// @Param json satellites
// @Success 200 {object} models.Ship
// @Router /location/topsecret [get]
func (h locationHandlers) GetLocationBySatellites() echo.HandlerFunc {
	return func(c echo.Context) error {
		var shipRequest models.Request
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "locationHandlers.GetLocationBySatellites")
		defer span.Finish()
		r := c.Request()
		err := json.NewDecoder(r.Body).Decode(&shipRequest)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		shipPositionAndMessage, err := h.locationUC.GetLocationBySatellites(ctx, shipRequest)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, shipPositionAndMessage)
	}
}

func (h locationHandlers) PostTopSecretSplit() echo.HandlerFunc {

	return nil
}

func (h locationHandlers) GetTopSecretSplit() echo.HandlerFunc {

	return nil
}
