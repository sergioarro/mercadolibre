package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
	"mercadolibre/pkg/httpErrors"
	"mercadolibre/pkg/logger"
)

const (
	basePrefix    = "api-location:"
	cacheDuration = 3600
)

type locationUC struct {
	cfg          *config.Config
	locationRepo location.Repository
	logger       logger.Logger
}

// Location UseCase constructor
func NewLocationUseCase(cfg *config.Config, locationRepo location.Repository, logger logger.Logger) location.UseCase {
	return &locationUC{cfg: cfg, locationRepo: locationRepo, logger: logger}
}

func (u *locationUC) GetLocationBySatellites(ctx context.Context, satellites models.Request) (*models.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationUC.GetLocationBySatellites")
	defer span.Finish()

	if len(satellites.RequestSatellites) > 3 {
		err := errors.New("Too many satellities")
		return nil, httpErrors.NewRestError(http.StatusPreconditionRequired, "Too many satellities", errors.Wrap(err, "locationUC.GetLocationBySatellites.ValidToManySatellities"))
	}

	position := models.Position{X: -100.0, Y: 75.5}
	n := &models.Response{
		Position: position,
		Message:  "este es un mensaje secreto",
	}

	/*n, err := u.locationRepo.GetLocationBySatellites(ctx, satellites)
	if err != nil {
		return nil, err
	}*/

	return n, nil
}

func (u *locationUC) getKeyWithPrefix(newsID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, newsID)
}

//get all the satellites that are operating
func (u *locationUC) getAllSatellitesInService() (SatellitesInService []models.Satellite) {
	SatellitesInService = append(SatellitesInService,
		models.Satellite{Name: "Kenobi", Position: models.Position{X: -500, Y: -200}},
		models.Satellite{Name: "Skywalker", Position: models.Position{X: 100, Y: -100}},
		models.Satellite{Name: "Sato", Position: models.Position{X: 500, Y: 100}})
	return
}
